package main

import "testing"

var str string

func BenchmarkStringBuilderNative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = buildStrNaive()
	}
}

func BenchmarkStringBuilderBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		str = buildStrBuilder()
	}
}

// go test -bench=. -benchmem
