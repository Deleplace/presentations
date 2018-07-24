package main

import (
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	count = 0
	for i := 0; i < 200; i++ {
		go h(nil, nil)
	}

	time.Sleep(3 * time.Second)
	if count != 200 {
		t.Errorf("Expected %d hits, got %d", 200, count)
	}
}
