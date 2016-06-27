package main

import "net/http"

func main() {
	http.HandleFunc("/", h)
	http.ListenAndServe(":7070", nil)
}

var count int = 0

func h(w http.ResponseWriter, r *http.Request) {
	count++
}
