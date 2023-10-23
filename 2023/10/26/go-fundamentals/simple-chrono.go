package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	t := time.Now() // HL

	fetchData()

	duration := time.Since(t) // HL

	fmt.Println("Took", duration)
	// END OMIT
}

func fetchData() {
	fmt.Println("Fetching data...")
	// Simulate a 1500ms roundtrip
	time.Sleep(1500 * time.Millisecond)
}
