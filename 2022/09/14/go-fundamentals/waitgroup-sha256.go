package main

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

func main() {
	// START OMIT
	inputs := []string{
		"hello world",
		"hello India",
		"everything is awesome",
	}

	fmt.Println("STARTING")
	var wg sync.WaitGroup // HL
	wg.Add(len(inputs))   // HL
	for i := range inputs {
		s := inputs[i]
		go func() {
			h := sha256.New()
			h.Write([]byte(s))
			fmt.Printf("%x\n", h.Sum(nil))
			wg.Done() // HL
		}()
	}
	wg.Wait() // HL
	fmt.Println("DONE")
	// END OMIT
}
