package main

import "fmt"

func main() {
	// START OMIT
	ch := make(chan string)

	go func() {
		msg := <-ch
		fmt.Println(msg)
		ch <- "pong"
	}()

	ch <- "ping"
	msg := <-ch
	fmt.Println(msg)
	close(ch)
	fmt.Println("THE END")
	// END OMIT
}
