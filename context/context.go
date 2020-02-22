package main

import (
	"context"
	"fmt"
	"time"
)

func contextExample() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exist, stop...")
				return
			default:
				fmt.Println("goroutine...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("notice stop")
	cancel()
	time.Sleep(5 * time.Second)
}

func contextMultiRoutine() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "observe 1")
	go watch(ctx, "observe 2")
	go watch(ctx, "observe 3")

	time.Sleep(10 * time.Second)
	fmt.Println("notice stop")
	cancel()
	time.Sleep(5 * time.Second)
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exist, stop...")
			return
		default:
			fmt.Println("goroutine...")
			time.Sleep(2 * time.Second)
		}
	}
}
