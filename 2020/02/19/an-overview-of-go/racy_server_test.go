package main

import (
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	count = 0
	go main()
	for i := 0; i < 200; i++ {
		go http.Get("http://localhost:7070/")
	}

	time.Sleep(3 * time.Second)
	if count != 200 {
		t.Errorf("Expected %d hits, got %d", 200, count)
	}
}
