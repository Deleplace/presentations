package main

import (
	"fmt"
	"regexp"
)

// START OMIT
func main() {
	fmt.Println("Hello")
	run()
	fmt.Println("Good bye")
}

func run() {
	defer func() {
		msg := recover() // HL
		fmt.Println("Oops:", msg)
	}()

	r := regexp.MustCompile("\\") // HL
	fmt.Println(r.MatchString("c:\\"))
}

// END OMIT
