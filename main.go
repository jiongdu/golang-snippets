package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	inFile := "./data/country.txt"

	fi, err := os.Open(inFile)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	index := 0
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		index++
		country := string(a)
		res := "\"" + country + "\"" + " : " + "{},"
		if index == 1 {
			fmt.Println(res)
		}
	}
	fmt.Println()
	fmt.Println("total line is ", index)

}

