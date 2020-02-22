package main

import (
	"sync"
	"testing"
)

func TestAppend(t *testing.T) {
	// no data race, because there is no memory to place, so it makes new memory
	x := []string{"start"}
	// data race, two goroutines notice there is memory to place, race happends because both goroutines trying to write the same memory
	// x := make([]string, 0, 6)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		y := append(x, "hello", "world")
		t.Log(cap(y), len(y))
	}()

	go func() {
		defer wg.Done()
		y := append(x, "goodbye", "bob")
		t.Log(cap(y), len(y))
	}()
	wg.Wait()
	t.Log(x)
}
