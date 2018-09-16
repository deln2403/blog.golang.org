// Passing slices to functions
// iterating over the indices of a slice (using a for range loop), incrementing its elements
// Even though the slice header is passed by value, the header includes a pointer to elements of an array,
// so both the original slice header and the copy of the header passed to the function describe the same array.
// Therefore, when the function returns, the modified elements can be seen through the original slice variable.
package main

import (
	"fmt"
)

var buffer [256]byte

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func main() {
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
}
