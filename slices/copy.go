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

func main() {
	slice := make([]int, 10, 15)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

	newSlice := make([]int, len(slice), 2*cap(slice))
	copy(newSlice, slice)
	fmt.Printf("len: %d, cap: %d\n", len(newSlice), cap(newSlice))
}
