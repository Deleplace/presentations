package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"
)

const url = "http://localhost:7070/"
const (
	N = 100
	M = 150
)

func main() {
	clock := func(f func()) time.Duration {
		t := time.Now()
		f()
		return time.Since(t)
	}
	d := clock(inject)
	fmt.Println(M*N, "requests processed in", d)
}

func inject() {
	for k := 0; k < M; k++ {
		var wg sync.WaitGroup
		wg.Add(N)
		for i := 0; i < N; i++ {
			go func() {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				} else {
					io.Copy(ioutil.Discard, resp.Body)
					resp.Body.Close()
				}
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
