// Some notes:
// 1. Methods are orthogonal to the underlying data.
// 2. One might similarly ask  "Why can I pass a nil pointer
//    to a function?" The answer is: because  pointers can be nil
// 3. Reusing similar pieces of code.

package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) sayHello() {
	if p == nil {
		fmt.Println("Who I am?")
		return
	}
	fmt.Println("Hello, my name is:", p.first, p.last)
}

func main() {
	var p *person
	p.sayHello()
}
