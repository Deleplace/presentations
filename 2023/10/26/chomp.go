package main

import "fmt"

func main() {
	chomp := func(s string) string { return s[:len(s)-1] }

	t := chomp("/usr/local/")

	fmt.Println(t)
}
