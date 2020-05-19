package zk_util

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

type ZkCallBackFunc func(ctx *context.Context, data []byte) interface{}

type ZkDataUtil struct {
	zkDataMutex   sync.Mutex     //zkData Mutex，用来互斥从zk中获得的数据
	zkPath        string         //zk 节点 路径
	zkStat        *zk.Stat       //zk 节点的状态
	zkData        []byte         //从zk 读取的数据
	zkConn        *zk.Conn       //zk连接
	zkErr         error          //getW回调 以及 定时器出错的error
	zkErrMutex    sync.Mutex     //error Mutex,用于error信息的互斥修改
	callBackFunc  ZkCallBackFunc //用户自定义的回调函数
	processedData interface{}    //用户自定义的回调函数处理的结果数据
}

const (
	ZkTimeTick       = time.Minute * 1 //zk 定时器 定时设置时间
	zkConnectTryTime = 3               //zk 尝试连接时间
)

func (zkData *ZkDataUtil) Init(ctx *context.Context, zkHost string, zkPath string, setCallBack bool, callBackFunc ZkCallBackFunc) (err error) {
	zkData.zkPath = zkPath

	for i := 0; i < zkConnectTryTime; i++ {
		zkData.zkConn, _, err = zk.Connect([]string{zkHost}, time.Second*5)
		if err == nil {
			break
		}
	}
	if err != nil {
		log.Fatal(err)
		return err
	}
	if setCallBack {
		zkData.callBackFunc = callBackFunc
		// 先执行一次获取数据操作
		zkData.GetDataFromZKWithCallBack(ctx)
		go func() {
			//?
			zkData.ZkGetByTimeWithCallBack(ctx)
		}()
		go func() {
			zkData.ZkGetWatcherWithCallBack(ctx)
		}()
	} else {
		// 先执行一次获取数据操作
		zkData.GetDataFromZK(ctx)
		go func() {
			//开启协程进行定时器任务
			zkData.ZkGetByTime(ctx)
		}()
		go func() {
			//开启协程进行zk回调任务
			zkData.ZkGetWatcher(ctx)
		}()
	}
	return nil
}

func (zkData *ZkDataUtil) GetDataFromZK(ctx *context.Context) {
	data, _, err := zkData.zkConn.Get(zkData.zkPath)
	if err != nil {
		zkData.zkErrMutex.Lock()
		zkData.zkErr = err
		zkData.zkErrMutex.Unlock()
	} else {
		zkData.zkErrMutex.Lock()
		zkData.zkErr = nil
		zkData.zkErrMutex.Unlock()
		//互斥修改zkData数据
		zkData.zkDataMutex.Lock()
		zkData.zkData = data
		zkData.zkDataMutex.Unlock()
	}
}

//从zk节点获取data数据，并使用用户自定义回调函数进行数据处理，获得数据存入zkData结构体的processedData结构体
func (zkData *ZkDataUtil) GetDataFromZKWithCallBack(ctx *context.Context) {
	data, _, err := zkData.zkConn.Get(zkData.zkPath)
	if err != nil {
		//互斥修改zkErr数据
		zkData.zkErrMutex.Lock()
		zkData.zkErr = err
		zkData.zkErrMutex.Unlock()
	} else {
		//互斥修改zkErr数据
		zkData.zkErrMutex.Lock()
		zkData.zkErr = nil
		zkData.zkErrMutex.Unlock()

		//调用用户自定义回调函数处理zk节点数据，并互斥修改 processedData
		processedData := zkData.callBackFunc(ctx, data)
		zkData.zkDataMutex.Lock()
		zkData.processedData = processedData
		zkData.zkDataMutex.Unlock()
	}
}

//带有用户自定义数据处理回调函数的zk GetW 协程处理函数
func (zkData *ZkDataUtil) ZkGetWatcherWithCallBack(ctx *context.Context) {
	for {
		//从zk获取数据，并添加一个watcher
		data, _, ch, err := zkData.zkConn.GetW(zkData.zkPath)
		if err != nil {
			//互斥修改zkErr数据
			zkData.zkErrMutex.Lock()
			zkData.zkErr = err
			zkData.zkErrMutex.Unlock()
		}

		//互斥修改zkErr数据
		zkData.zkErrMutex.Lock()
		zkData.zkErr = nil
		zkData.zkErrMutex.Unlock()

		//调用用户自定义回调函数处理zk节点数据，并互斥修改 processedData
		processedData := zkData.callBackFunc(ctx, data)
		zkData.zkDataMutex.Lock()
		zkData.processedData = processedData
		zkData.zkDataMutex.Unlock()

		//阻塞等待zk回调事件触发
		select {
		case ev := <-ch:
			if ev.Err != nil {
				//互斥修改zkErr数据
				zkData.zkErrMutex.Lock()
				zkData.zkErr = ev.Err
				zkData.zkErrMutex.Unlock()
			}
			if ev.Path != zkData.zkPath {
				pathErr := errors.New("GetW watcher wrong path")
				//互斥修改zkErr数据
				zkData.zkErrMutex.Lock()
				zkData.zkErr = pathErr
				zkData.zkErrMutex.Unlock()
			}
		}
	}
}

func (zkData *ZkDataUtil) ZkGetWatcher(ctx *context.Context) {

	for {
		//从zk获取数据，并添加一个watcher
		data, stat, ch, err := zkData.zkConn.GetW(zkData.zkPath)
		if err != nil {
			//互斥修改zkErr数据
			zkData.zkErrMutex.Lock()
			zkData.zkErr = err
			zkData.zkErrMutex.Unlock()
		}

		//互斥修改zkErr数据
		zkData.zkErrMutex.Lock()
		zkData.zkErr = nil
		zkData.zkErrMutex.Unlock()

		//互斥修改zkData数据
		zkData.zkDataMutex.Lock()
		zkData.zkData = data
		zkData.zkStat = stat
		zkData.zkDataMutex.Unlock()

		//阻塞等待zk回调事件触发
		select {
		case ev := <-ch:
			if ev.Err != nil {
				//互斥修改zkErr数据
				zkData.zkErrMutex.Lock()
				zkData.zkErr = ev.Err
				zkData.zkErrMutex.Unlock()
			}
			if ev.Path != zkData.zkPath {
				pathErr := errors.New("GetW watcher wrong path")
				//互斥修改zkErr数据
				zkData.zkErrMutex.Lock()
				zkData.zkErr = pathErr
				zkData.zkErrMutex.Unlock()
			}
		}
	}
}

func (zkData *ZkDataUtil) ZkGetByTimeWithCallBack(ctx *context.Context) {
	zkTicker := time.NewTicker(ZkTimeTick)
	for {
		select {
		case <-zkTicker.C:
			zkData.GetDataFromZKWithCallBack(ctx)
		}
	}
}

func (zkData *ZkDataUtil) ZkGetByTime(ctx *context.Context) {
	zkTicker := time.NewTicker(ZkTimeTick)
	for {
		select {
		case <-zkTicker.C:
			zkData.GetDataFromZK(ctx)
		}
	}
}

func (zkData *ZkDataUtil) GetDataWithCallback() (interface{}, error) {
	var data interface{}
	var err error
	zkData.zkErrMutex.Lock()
	err = zkData.zkErr
	zkData.zkErrMutex.Unlock()

	zkData.zkDataMutex.Lock()
	data = zkData.processedData
	zkData.zkDataMutex.Unlock()
	return data, err
}

func (zkData *ZkDataUtil) GetData() (string, error) {
	var data []byte
	var err error

	zkData.zkErrMutex.Lock()
	err = zkData.zkErr
	zkData.zkErrMutex.Unlock()

	zkData.zkDataMutex.Lock()
	data = zkData.zkData
	zkData.zkDataMutex.Unlock()

	return string(data), err
}
