package inlining

import (
	"fmt"
	"os"
	"testing"
)

func f1(n int) int {
	if n%242424242 == 242424241 {
		fmt.Fprintf(os.Stderr, "Found %d\n", n)
	}
	return n * n
}

func BenchmarkF1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = f1(i)
	}
}

//go:noinline
func f2(n int) int {
	if n%242424242 == 242424241 {
		fmt.Fprintf(os.Stderr, "Found %d\n", n)
	}
	return n * n
}

func BenchmarkF2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sink = f2(i)
	}
}

// Sink prevents the compiler from optimizing away
// an unused result
var Sink int
