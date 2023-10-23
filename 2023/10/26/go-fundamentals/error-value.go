package main

import (
	"fmt"
	"regexp"
)

// START OMIT
func main() {
	ok, err := matches("a", "abcde")
	fmt.Println(ok, err)
	ok, err = matches("(", "abcde")
	fmt.Println(ok, err)
	fmt.Println("THE END")
}

func matches(pattern string, s string) (bool, error) { // HL
	r, err := regexp.Compile(pattern)
	if err != nil {
		return false, fmt.Errorf("matches failed: %v", err)
	}
	return r.MatchString(s), nil
}

// END OMIT
