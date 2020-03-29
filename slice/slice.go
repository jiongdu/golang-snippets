package main

import "fmt"

func a() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	y := append(x, 2)
	z := append(x, 3)
	fmt.Println(y, z)
}

func b() {
	x := []int{}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2)
	y := append(x, 3)
	z := append(x, 4)
	fmt.Println(y, z)
}

func c() {
	x := []int{1, 2, 3}
	y := make([][]int, 0, 5)
	y = append(y, x)
	z := make([][]int, 0, 5)
	z = append(z, append([]int{}, x...))
	// fmt.Println(x, y)
	// x = append(x, 4)
	fmt.Println(x, y, z)
	x[0] = 9
	fmt.Println(x, y, z)
}

func insert(s *[][]int, p int, v []int) {
	*s = append(*s, []int{})
	copy((*s)[p+1:], (*s)[p:])
	(*s)[p] = v
}

// slice is reference type
func main() {

	//一个理解slice工作原理的好例子
	// a()
	// b()

	c()

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
