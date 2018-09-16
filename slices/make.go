package main

import (
	"fmt"
)

func main() {
	//  creates a slice of length 10 with room for 5 more (15-10)
	slice := make([]int, 10, 15)
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))

	// doubles the capacity of our int slice but keeps its length the same
	newSlice := make([]int, len(slice), 2*cap(slice))
	for i := range slice {
		newSlice[i] = slice[i]
	}
	slice = newSlice
	fmt.Printf("len: %d, cap: %d\n", len(slice), cap(slice))
}
