package main

import (
	"fmt"
	"regexp"
)

var r = regexp.MustCompile("\\") // HL

func main() {
	fmt.Println("Hello")
	fmt.Println(r.MatchString("c:\\"))
	fmt.Println("Good bye")
}
