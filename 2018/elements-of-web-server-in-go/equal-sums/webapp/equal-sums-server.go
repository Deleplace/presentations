package main

// Solution for problem "Equal Sums" from Google Code Jam 2012
//
// See https://code.google.com/codejam/contest/1836486/dashboard#s=p2

import (
	//	"appengine"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
	//	"runtime"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		paramCase := r.FormValue("case")
		caseStr := strings.Split(paramCase, ",")
		caseInts, err := string2int64(caseStr)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		//caseInts := []int64{120, 266, 858, 1243, 1657, 1771, 2328, 2490, 2665, 2894, 3117, 4210, 4454, 4943, 5690, 6170, 7048, 7125, 9512, 9600}
		msg := solve(caseInts)

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("(Go) \n"))
		w.Write([]byte(msg))
	})

	port := ":8080"
	fmt.Println("Starting server on", port)
	err := http.ListenAndServe(port, nil)
	log.Fatal(err)
}

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
		s := sum(cur)
		if other, ok := memo[s]; ok {
			sort.Sort(&other)
			sort.Sort(&cur)
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

func sum(x sol) int64 {
	s := int64(0)
	for _, a := range x {
		s += a
	}
	return s
}

func (x *sol) Len() int           { return Y }
func (x *sol) Less(i, j int) bool { return x[i] < x[j] }
func (x *sol) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

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
