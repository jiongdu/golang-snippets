package main

import (
	"sync"
	"testing"
)

func TestAppend(t *testing.T) {
	x := []string{"start"} // no data race
	// x := make([]string, 0, 6) // data race

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
