package main

import (
	"fmt"
	"sync"
)

// START OMIT
type GuardedString struct {
	sync.Mutex // HL

	S string
}

func main() {
	s := GuardedString{S: "important data"}
	fmt.Println(s)

	s.Lock()
	s.S = "altered data"
	s.Unlock()
	fmt.Println(s)
}

// END OMIT

func (gs GuardedString) String() string {
	return fmt.Sprintf("Guarded String: %q", gs.S)
}
