package main

import (
	"fmt"

	"./pkg"
)

func main() {
	fmt.Println(pkg.Boring)
}

// Notes:
// 1. Useful when you're working with people that don't have
//    Go environemt installed, and you still want to split
//    your project into packages. For example, submitting your
//    University HW.
// 2. Doesn't work well with vendor. see: #15478.
