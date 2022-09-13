package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

func main() {
	// START OMIT
	s := "SGVsbG8gSW5kaWE="

	r := strings.NewReader(s)
	d := base64.NewDecoder(base64.StdEncoding, r) // HL
	t, _ := io.ReadAll(d)
	fmt.Println(string(t))
	// END OMIT
}
