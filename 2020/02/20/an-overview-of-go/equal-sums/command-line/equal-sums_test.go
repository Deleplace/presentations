package main

import (
	"math/rand"
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := randomInput()
		solution := solve(in)
		_ = solution
	}
}

func BenchmarkSolveTwice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(randomInput())
		solve(randomInput())
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
