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

func f2part1(n int) int {
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
	return x
}

func f2part2(x int) int {
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
		Sink = f2part2(f2part1(i))
	}
}

func TestSameResults(t *testing.T) {
	for i := 1; i < 10000; i++ {
		if a, b := f1(i), f2part2(f2part1(i)); a != b {
			i := i
			t.Fatalf("f1(%d)==%d != f2(%d)==%d", i, a, i, b)
		}
	}
}

// Sink prevents the compiler from optimizing away
// an unused result
var Sink int
