/*
When we doubled the capacity of our slice in the previous section (make.go),
we wrote a loop to copy the old data to the new slice. Go has a built-in function,
copy, to make this easier. Its arguments are two slices, and it copies the data
from the right-hand argument to the left-hand argument
*/

package main

import (
	"fmt"
)

// Insert inserts the value into the slice at the specified index,
// which must be in range.
// The slice must have room for the new element.
func Insert(slice []int, index, value int) []int {
	// Grow the slice by one element.
	slice = slice[0 : len(slice)+1]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	// Store the new value.
	slice[index] = value
	// Return the result.
	return slice
}

func main() {
	slice := make([]int, 10, 15)

	// Here's how to use copy to insert a value into the middle of a slice.
	for i := range slice {
		slice[i] = i
	}
	fmt.Println(slice)
	slice = Insert(slice, 5, 99)
	fmt.Println(slice)
}
