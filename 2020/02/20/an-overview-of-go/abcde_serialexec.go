package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// START OMIT
	A()
	B()
	C()
	D()
	E()
	// END OMIT
}

func A() {
	fmt.Println("Starting A")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   A")
}

func B() {
	fmt.Println("Starting B")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   B")
}

func C() {
	fmt.Println("Starting C")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   C")
}

func D() {
	fmt.Println("Starting D")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   D")
}

func E() {
	fmt.Println("Starting E")
	time.Sleep(time.Duration(rand.Intn(10)) * 100 * time.Millisecond)
	fmt.Println("End of   E")
}
