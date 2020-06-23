package dataloader

import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
)

func main() {
	ms := make(map[string]int)
	ms["key1"] = 1
	ms["key2"] = 2

	batchFunc := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		var results []*dataloader.Result
		results = append(results, &dataloader.Result{
			//Data:  m[keys.Keys()[0]],
			Error: nil,
		})
		return results
	}

	loader := dataloader.NewBatchedLoader(batchFunc)

	thunk := loader.Load(context.TODO(), dataloader.StringKey("key2"))
	result, err := thunk()
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}
	fmt.Println("value:", result)
}
