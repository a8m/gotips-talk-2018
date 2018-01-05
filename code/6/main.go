package main

import (
	"fmt"

	"./pkg"
)

func main() {
	fmt.Println(pkg.Boring)
}

// Notes:
// 1. Doesn't work well with vendor. see: #15478.
