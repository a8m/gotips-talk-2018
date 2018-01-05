// An example to how control parallelism using a semaphore.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	var (
		jobs = 100
		s    = newsem(runtime.NumCPU())
	)
	for i := 0; i < jobs; i++ {
		s.acquire()
		go work(i, s.release)
	}
	s.wait()
	fmt.Println("done")
}

func work(i int, done func()) {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	fmt.Printf("finish work %d\n", i)
	done()
}
