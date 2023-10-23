package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) { // HL
	fmt.Fprintln(w, "Hello Deezer") // HL
} // HL

func main() {
	http.HandleFunc("/", handler)            // HL
	err := http.ListenAndServe(":8080", nil) // HL
	log.Fatal(err)
}
