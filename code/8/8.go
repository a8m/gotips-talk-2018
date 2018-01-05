package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	quit := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 1e4; i++ {
		wg.Add(1)
		go boring(wg.Done, quit)
	}
	time.Sleep(time.Second)
	// notify.
	close(quit)
	// join.
	wg.Wait()

	// Notes:
	// 1. could also done with context (it does the same, btw).
	// 2. explain the gotcha.
}

func boring(done func(), quit <-chan struct{}) {
	quit = nil
	defer done()
		r := rand.Intn(1e3)
		select {
		case <-time.After(time.Duration(r) * time.Millisecond):
			fmt.Println("boring", r)
		case <-quit:
			time.Sleep(time.Duration(r))
			return
		}
	}
}
