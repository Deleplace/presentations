package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Fetching data...")

	// START OMIT
	duration := clock(func() {
		for _, r := range resources {
			fetch(r)
		}
	})
	// END OMIT

	fmt.Println("Took", duration)
}

func clock(f func()) time.Duration {
	t := time.Now()
	f()
	return time.Since(t)
}

type Resource struct {
	URL string
}

var resources = []Resource{
	{URL: "https://en.wikipedia.org/wiki/Apple"},
	{URL: "https://en.wikipedia.org/wiki/Banana"},
	{URL: "https://en.wikipedia.org/wiki/Cherry"},
}

func fetch(r Resource) {
	// Simulate a random roundtrip
	d := time.Duration(100+rand.Intn(800)) * time.Millisecond
	time.Sleep(d)
}
