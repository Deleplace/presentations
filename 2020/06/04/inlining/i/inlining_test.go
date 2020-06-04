package inlining

import (
	"math/rand"
	"testing"
)

func grey1(c []byte) byte {
	var sum int
	for _, v := range c[:3] {
		sum += int(v)
	}
	avg := sum / 3
	return byte(avg)
}

func BenchmarkGrey1(b *testing.B) {
	c := randomColor()
	for i := 0; i < b.N; i++ {
		Sink = grey1(c)
	}
}

func grey2(c []byte) byte {
	var sum int
	sum += int(c[0])
	sum += int(c[1])
	sum += int(c[2])
	avg := sum / 3
	return byte(avg)
}

func BenchmarkGrey2(b *testing.B) {
	c := randomColor()
	for i := 0; i < b.N; i++ {
		Sink = grey2(c)
	}
}

func grey3(c []byte) byte {
	var sum int
	sum += int(c[2])
	sum += int(c[1])
	sum += int(c[0])
	avg := sum / 3
	return byte(avg)
}

func BenchmarkGrey3(b *testing.B) {
	c := randomColor()
	for i := 0; i < b.N; i++ {
		Sink = grey3(c)
	}
}

func randomColor() []byte {
	return []byte{byte(rand.Intn(256)), byte(rand.Intn(256)), byte(rand.Intn(256))}
}

var Sink byte
