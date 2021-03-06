package main

import (
	"fmt"
	"sync"
)

// we can use "go run -race data_race.go" to detect data race condition
// the method to solve data race: WaitGroup, channel, mutex
func main() {
	fmt.Println(getNumber())

	var m sync.Map
}

func getNumber() int {
	// var wg sync.WaitGroup
	// wg.Add(1)
	var i int
	go func() {
		i = 5
		// wg.Done()
	}()
	// wg.Wait()
	return i
}
