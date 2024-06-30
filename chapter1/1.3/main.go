package main

import (
	"strings"
	"testing"
)

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		input := []string{"12", "qw", "as", "zx"}
		strings.Join(input, " ")
	}
}
	