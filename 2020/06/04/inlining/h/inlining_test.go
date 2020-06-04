package inlining

import (
	"testing"
)

func f1(m int) {
	if m == 42 {
		if m < Sink*Sink*Sink*Sink*Sink {
			println("You win")
		}
	}
}

func BenchmarkF1(b *testing.B) {
	m := 13
	for i := 0; i < b.N; i++ {
		f1(m)
	}
}

//go:noinline
func f2(m int) {
	if m == 42 {
		if m < Sink*Sink*Sink*Sink*Sink {
			println("You win")
		}
	}
}

func BenchmarkF2(b *testing.B) {
	m := 13
	for i := 0; i < b.N; i++ {
		f2(m)
	}
}

func BenchmarkF1bis(b *testing.B) {
	m := 42
	for i := 0; i < b.N; i++ {
		f1(m)
	}
}

func BenchmarkF2bis(b *testing.B) {
	m := 42
	for i := 0; i < b.N; i++ {
		f2(m)
	}
}

func BenchmarkNoOp(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

var Sink int
