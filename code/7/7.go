// Run the vet command like this the "-shadow" flag.
package main

import "os"

func main() {

}

func BadRead(f *os.File, buf []byte) error {
	var err error
	for {
		n, err := f.Read(buf) // shadows the function variable 'err'
		if err != nil {
			break // causes return of wrong value
		}
		foo(buf)
	}
	return err
}
