// What's the problem here?
package main

import (
	"fmt"
	"math/rand"
	"runtime"

	"code.cloudfoundry.org/bytefmt"
)

const size = 10

func boring() []byte {
	// get blob from some stream.
	blob := make([]byte, 1e8)
	// "search something", and do slicing.
	n := rand.Intn(size)
	return blob[n : n+size]
}

func main() {
	alloc()
	buf := boring()
	// run the garbage collector
	runtime.GC()
	alloc()
	fmt.Println("buf:", buf)
}

func alloc() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Println("n bytes of allocated heap objects:", bytefmt.ByteSize(m.Alloc))
}
