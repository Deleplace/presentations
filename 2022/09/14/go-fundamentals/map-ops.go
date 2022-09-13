package main

import "fmt"

func main() {
	// START OMIT
	ops := map[string]func(int, int) int{"+": plus, "-": minus, "*": multiply, "/": divide}

	a, b := 13, 5

	for symbol, op := range ops {
		fmt.Println(a, symbol, b, "=", op(a, b))
	}
	// END OMIT
}

func plus(x, y int) int     { return x + y }
func minus(x, y int) int    { return x - y }
func multiply(x, y int) int { return x * y }
func divide(x, y int) int   { return x / y }
