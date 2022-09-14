package main

import (
	"math/rand"
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	in := randomInput()
	for i := 0; i < b.N; i++ {
		solution := solve(in)
		Sink = solution
	}
}

func randomInput() []int64 {
	const N = 50
	in := make([]int64, N)
	for i := range in {
		in[i] = int64(1 + rand.Intn(1000))
	}
	return in
}

// Sink prevents optimizing computation away
var Sink string
