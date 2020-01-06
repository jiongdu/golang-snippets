package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix()) //unix timestamp
	fmt.Println(time.Now().Day())
	t := time.Now()
	curZeroTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	// fmt.Println(curZeroTime.Date())
	fmt.Println(curZeroTime)
	curZeroTimeUnix := curZeroTime.Unix()
	ysdZeroTimeUnix := curZeroTimeUnix - 86400

	fmt.Println(curZeroTimeUnix)
	fmt.Println(ysdZeroTimeUnix)
}
