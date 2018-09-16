// Pointers to slices: Method receivers
package main

import (
	"bytes"
	"fmt"
)

var buffer [256]byte

type path []byte

func (p *path) TruncateAtFinalSlash() {
	i := bytes.LastIndex(*p, []byte("/"))
	if i >= 0 {
		*p = (*p)[0:i]
	}
}

func main() {
	// method on a slice that truncates it at the final slash
	initPath := "/usr/bin/tso"
	pathName := path(initPath) // Conversion from string to path.
	pathName.TruncateAtFinalSlash()
	fmt.Printf("%s -> %s\n", initPath, pathName)
}
