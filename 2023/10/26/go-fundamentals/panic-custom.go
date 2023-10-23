package main

import "fmt"

func main() {
	a := -9
	// START OMIT
	switch a % 2 {
	case 0:
		fmt.Println("even")
	case 1:
		fmt.Println("odd")
	default:
		panic(a % 2) // HL
	}
	// END OMIT
}
