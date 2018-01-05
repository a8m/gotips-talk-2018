package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{1, 5, 3, 4, 2, 7, 6}
	sort.Sort(Reverse(sort.IntSlice(s)))
	fmt.Println(s)
}

// Code taken from golang/pkg/sort
type reverse struct {
	// This embedded Interface permits Reverse to use the methods of
	// another Interface implementation.
	sort.Interface
}

// Here we can override a specific method without having to define all the
// others.
//
// Less returns the opposite of the embedded implementation's Less method.
func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// Reverse returns the reverse order for data.
func Reverse(data sort.Interface) sort.Interface {
	return &reverse{data}
}
