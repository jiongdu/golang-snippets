/*
 2 producer and 1 consumer
 produce1() publish "x", "y" to topic "topic_test"
 produce2() publish "z" to topic "topic_test"
 consume1() subscribe channel "channel_sensor01" of topic "topic_test"
*/

package example

import (
	"log"
	"testing"
	"time"

	"github.com/nsqio/go-nsq"
)

func TestNSQ1(t *testing.T) {
	NSQsAddrs := []string{"127.0.0.1:4150", "127.0.0.1:4152"}
	go consume1(NSQsAddrs)
	go produce1()
	go produce2()
	time.Sleep(30 * time.Second)
}

func produce1() {
	cfg := nsq.NewConfig()
	nsqdAddr := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(nsqdAddr, cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	if err := producer.Publish("topic_test", []byte("x")); err != nil {
		log.Fatal("publish error:" + err.Error())
	}
	if err := producer.Publish("topic_test", []byte("y")); err != nil {
		log.Fatal("publish error:" + err.Error())
	}
}

func produce2() {
	cfg := nsq.NewConfig()
	nsqdAddr := "127.0.0.1:4152"
	producer, err := nsq.NewProducer(nsqdAddr, cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := producer.Publish("topic_test", []byte("z")); err != nil {
		log.Fatal("publish error:" + err.Error())
	}
}

func consume1(NSQsAddrs []string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("topic_test", "channel_sensor01", cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body) + "Consumer1")
		return nil
	}))
	//这里是直连，还可以通过查询nsdlookupd来发现指定topic的生产者
	if err := consumer.ConnectToNSQDs(NSQsAddrs); err != nil {
		log.Fatal("err", " C1")
	}
	<-consumer.StopChan
}
