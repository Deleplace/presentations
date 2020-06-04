package inlining

import "testing"

func square1(n int) int {
	return n * n
}

func double1(n int) int {
	return n + n
}

func doubleSquare1(n int) int {
	x := square1(n)
	x = double1(x)
	return x
}

func BenchmarkDoubleSquare1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = doubleSquare1(i)
	}
}

//go:noinline
func square2(n int) int {
	return n * n
}

//go:noinline
func double2(n int) int {
	return n + n
}

//go:noinline
func doubleSquare2(n int) int {
	x := square2(n)
	x = double2(x)
	return x
}

func BenchmarkDoubleSquare2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = doubleSquare2(i)
	}
}

// Sink prevents the compiler from optimizing away
// an unused result
var Sink int
