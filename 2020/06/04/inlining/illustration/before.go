package main

import "fmt"

func square(n int) int {
	return n * n
}

func main() {
	x := square(9)
	fmt.Println(x)
}
