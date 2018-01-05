// popular example: http.HandlerFunc that implement the Handler interface.
//
// link: https://golang.org/src/net/http/server.go?s=57220:57267#L1928
package main

import (
	"fmt"
	"io"
	"strings"
)

// WriterFunc is a typed func that implements the io.Writer interface.
type WriterFunc func([]byte) (int, error)

func (fn WriterFunc) Write(p []byte) (int, error) { return fn(p) }

func main() {
	for _, writer := range []WriterFunc{
		// writer 1
		func(p []byte) (int, error) {
			fmt.Printf("writer-1 %s", string(p))
			return len(p), nil
		},
		// writer 2
		func(p []byte) (int, error) {
			fmt.Printf("writer-2 %s", string(p))
			return len(p), nil
		},
		// writer 3
		func(p []byte) (int, error) {
			fmt.Printf("writer-3 %s", string(p))
			return len(p), nil
		},
	} {
		io.Copy(writer, strings.NewReader("boring...\n"))
	}
}
