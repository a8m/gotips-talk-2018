// How can we solve it?
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
	buf := make([]byte, size)
	copy(buf, blob[n:n+size])
	return buf
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

// Notes:
// 1. Similar to strings - tell my story with gjson.
// 2. Few words about array and slicing -
//    slice references the original array, and as long as the slice
//    is kept around, the garbage collector can't release the backing array.
