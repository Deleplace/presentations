package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	duration := clock(fetchData) // HL

	fmt.Println("Took", duration)
}

func clock(f func()) time.Duration {
	t := time.Now()
	f()
	return time.Since(t)
}

// END OMIT

func fetchData() {
	fmt.Println("Fetching data...")
	// Simulate a 1500ms roundtrip
	time.Sleep(1500 * time.Millisecond)
}
