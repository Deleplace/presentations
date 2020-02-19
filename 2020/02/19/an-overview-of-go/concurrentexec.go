package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Fetching data...")
	duration := clock(fetchData)
	fmt.Println("Fetched data in", duration)
}

func fetchData() {
	// START OMIT
	RunConcurrent(
		profile,
		news,
		weather,
	)
	// END OMIT
}

func clock(f func()) time.Duration {
	t := time.Now()
	f()
	return time.Since(t)
}

func profile() {
	time.Sleep(2 * time.Second)
	fmt.Println(" Got profile")
}

func news() {
	time.Sleep(1 * time.Second)
	fmt.Println(" Got news")
}

func weather() {
	time.Sleep(3 * time.Second)
	fmt.Println(" Got weather")
}

// RunConcurrent launches funcs,
// and waits for their completion.
func RunConcurrent(funcs ...func()) {
	var wg sync.WaitGroup
	wg.Add(len(funcs))
	for _, f := range funcs {
		f := f
		go func() {
			f()
			wg.Done()
		}()
	}
	wg.Wait()
}
