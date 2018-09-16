// Pointers to slices: Method receivers
/*
method for path that upper-cases the ASCII letters in the path (parochially ignoring non-English names),
the method could be a value because the value receiver will still point to the same underlying array.
*/

package main

import (
	"fmt"
)

type path []byte

func (p path) ToUpper() {
	for i, b := range p {
		if 'a' <= b && b <= 'z' {
			p[i] = b + 'A' - 'a'
		}
	}
}

func main() {
	pathName := path("/usr/bin/tso")
	pathName.ToUpper()
	fmt.Printf("%s\n", pathName)
}
