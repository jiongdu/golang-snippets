package main

import "strings"

var strs = []string{
	"here's",
	"a",
	"some",
	"long",
	"list",
	"of",
	"strings",
	"for",
	"you",
}

func buildStrNaive() string {
	var res string

	for _, s := range strs {
		res += s
	}

	return res
}

func buildStrBuilder() string {
	b := strings.Builder{}
	b.Grow(60)
	for _, v := range strs {
		b.WriteString(v)
	}
	return b.String()
}
