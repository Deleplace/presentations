package main

// Solution for problem "Equal Sums" from Google Code Jam 2012
//
// See https://code.google.com/codejam/contest/1836486/dashboard#s=p2

import (
	"fmt"
	"os"

	"strconv"
)

func main() {
	if len(os.Args) < 9 {
		fmt.Fprintf(os.Stderr, "Usage: %s <list of 8+ integers> \n", os.Args[0])
		os.Exit(1)
	}

	caseStr := os.Args[1:]
	caseInts, err := string2int64(caseStr)
	if err != nil {
		panic(err)
	}
	//caseInts := []int64{120, 266, 858, 1243, 1657, 1771, 2328, 2490, 2665, 2894, 3117, 4210, 4454, 4943, 5690, 6170, 7048, 7125, 9512, 9600}
	msg := solve(caseInts)

	fmt.Println(msg)
}

func string2int64(str []string) ([]int64, error) {
	ints := make([]int64, len(str))
	for i, s := range str {
		var err error
		ints[i], err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func int642string(ints []int64) []string {
	str := make([]string, len(ints))
	for i, val := range ints {
		str[i] = fmt.Sprintf("%v", val)
	}
	return str
}
