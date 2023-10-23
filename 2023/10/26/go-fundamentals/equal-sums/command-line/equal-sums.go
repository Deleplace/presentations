package main

// Solution for problem "Equal Sums" from Google Code Jam 2012
//
// See https://code.google.com/codejam/contest/1836486/dashboard#s=p2

import (
	"fmt"
	"math/rand"

	"strings"

	"golang.org/x/exp/slices"
)

// Y is the size of the subsets that will be randomly picked
const Y = 6

type sol [Y]int64

func solve(in []int64) string {
	N := len(in)
	memo := make(map[int64]sol)
	var cur sol
	for {
		for i := 0; i < Y; i++ {
			j := i + 1 + rand.Intn(N-i-1)
			in[j], in[i] = in[i], in[j]
		}
		copy(cur[:], in[0:Y])
		s := cur.sum()
		if other, ok := memo[s]; ok {
			other.Normalize()
			cur.Normalize()
			if other != cur {
				prettyOther := strings.Join(int642string(other[:]), " + ")
				prettyCur := strings.Join(int642string(cur[:]), " + ")
				return fmt.Sprintf("%v = %v\n%v = %v", prettyOther, s, prettyCur, s)
			}
		}
		memo[s] = cur
	}
	return ""
}

func (x sol) sum() int64 {
	s := int64(0)
	for _, a := range x {
		s += a
	}
	return s
}

func (s *sol) Normalize() {
	// Sort the Y numbers in ascending order, so that 2 identical solutions can
	// be effectively compared
	slices.Sort(s[:])
}
