package main

import "fmt"

func main() {
	a := make([]int, 10)
	b := a[4:]
	fmt.Println("B is:", b)
	c := a[:1]
	c = append(c, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println("C is:", c)
	fmt.Println("Wait, and b?:", b)
	// see the explanation below
}

// Explanation:
//
// It's important to remember that all the slices share the
// same backing array. and appending to a one may change the
// other (if you didn't reached to the end).
//
// How do you fix it? use copy.
