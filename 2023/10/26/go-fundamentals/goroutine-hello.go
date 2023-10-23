package main

import (
	"fmt"
	"time"
)

func main() {

	// START OMIT
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Print("Hello ")
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(500 * time.Millisecond)
	for i := 0; i < 5; i++ {
		fmt.Print("world ")
		time.Sleep(time.Second)
	}
	// END OMIT

}
