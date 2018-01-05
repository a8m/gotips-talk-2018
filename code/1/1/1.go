// What this program prints?

package main

import "fmt"

type person struct {
	first string
	last  string
}

func main() {
	var p *person
	fmt.Println("Hello, my name is:", p.first, p.last)
}
