package inlining

import "testing"

func square1(n int) int {
	return n * n
}

func BenchmarkSquare1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = square1(i)
	}
}

//go:noinline
func square2(n int) int {
	return n * n
}

func BenchmarkSquare2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = square2(i)
	}
}

// Sink prevents the compiler from optimizing away
// an unused result
var Sink int
