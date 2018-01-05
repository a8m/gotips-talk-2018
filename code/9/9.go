package main

import "fmt"

func main() {
	fmt.Println(Parse("(()))"))
}

func Parse(s string) (err error) {
	defer func() {
		if e := recover(); e != nil {
			pe, ok := e.(ParseError)
			if !ok {
				panic(e)
			}
			err = pe
		}
	}()
	parse(s)
	return
}

// Notes:
//
// Useful pattern when "every step" is expectation.
// Some examples in the std:
// - json/decode.go
// - go/parser.go
// - text/template/parse.go
// - encoding/gob/decode.go "catchError"
//
// For more info, see: https://github.com/golang/go/wiki/PanicAndRecover.
//
// > By convention, no explicit panic() should be allowed to cross a package boundary.
// > Indicating error conditions to callers should be done by returning error value.
// > Within a package, however, especially if there are deeply nested calls to non-exported
// > functions, it can be useful (and improve readability) to use panic to indicate error
// > conditions which should be translated into error for the calling function.
// > Below is an admittedly contrived example of a way in which a nested function and an
// > exported function may interact via this panic-on-error relationship.
