package inlining

import "testing"

func f1(n int) int {
	x := n * n
	x /= 5
	x %= 130
	x /= 7
	x *= 42
	x >>= 2
	x += 9
	x %= 150
	x /= 3
	x %= 40
	x *= x
	x /= 11
	x *= 42
	x >>= 2
	x += 9
	x %= 150
	x /= 11
	x *= 42
	x >>= 2
	x += 9
	x %= 150
	x /= 11
	x *= 42
	x >>= 2
	x++
	return x
}

func BenchmarkF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = f1(i)
	}
}

//go:noinline
func f2(n int) int {
	x := n * n
	x /= 5
	x %= 130
	x /= 7
	x *= 42
	x >>= 2
	x += 9
	x %= 150
	x /= 3
	x %= 40
	x *= x
	x /= 11
	x *= 42
	x >>= 2
	x += 9
	x %= 150
	x /= 11
	x *= 42
	x >>= 2
	x += 9
	x %= 150
	x /= 11
	x *= 42
	x >>= 2
	x++
	return x
}

func BenchmarkF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = f2(i)
	}
}

// Sink prevents the compiler from optimizing away
// an unused result
var Sink int
