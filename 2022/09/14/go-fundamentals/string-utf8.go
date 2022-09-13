package main

import "fmt"

func main() {
	// START OMIT
	s := "He told me: \"मुझे चपाती खाना बहुत पसंद है\"."

	for _, c := range s {
		fmt.Printf("%c\n", c)
	}

	fmt.Println("s has", len(s), "bytes,", len([]rune(s)), "runes")

	t := s[0:7]
	fmt.Printf("t = %q", t)
	// END OMIT
}
