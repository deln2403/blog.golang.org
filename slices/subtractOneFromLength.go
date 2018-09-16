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

func SubtractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice)-1]
	return slice
}

func main() {
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}

	// argument to the function really is a copy
	fmt.Println("Before: len(slice) =", len(slice))
	newSlice := SubtractOneFromLength(slice)
	fmt.Println("After:  len(slice) =", len(slice))
	fmt.Println("After:  len(newSlice) =", len(newSlice))
}
