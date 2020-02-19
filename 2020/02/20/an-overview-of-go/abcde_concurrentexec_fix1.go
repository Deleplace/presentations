package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// START OMIT
	RunConcurrent(A, B, C, D, E)
	// END OMIT
}

// RunConcurrent launches funcs,
// and waits for their completion.
func RunConcurrent(funcs ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(funcs))
	for _, f := range funcs {
		go func(g func()) {
			g()
			wg.Done()
		}(f)
	}
	wg.Wait()
}

func A() {
	fmt.Println("Starting A")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   A")
}

func B() {
	fmt.Println("Starting B")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   B")
}

func C() {
	fmt.Println("Starting C")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   C")
}

func D() {
	fmt.Println("Starting D")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   D")
}

func E() {
	fmt.Println("Starting E")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   E")
}
