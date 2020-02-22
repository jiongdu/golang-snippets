package main

import "fmt"

// slice is reference type
func main() {
	s := make([]int, 0)

	for i := 0; i < 8; i++ {
		s = append(s, i)
	}

	//r := s
	r := make([]int, 4, 4)
	copy(r, s)
	r[0] = 8
	fmt.Println(s)
	fmt.Println(r)
}
