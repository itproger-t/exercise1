package main

import (
	"fmt"
)

func removeAtIndex(slice []string, i int) []string {
	if i < 0 || i >= len(slice) {
		return slice
	}

	newSlice := make([]string, len(slice)-1)
	copy(newSlice, slice[:i])
	copy(newSlice[i:], slice[i+1:])
	return newSlice
}

func main() {
	data := []string{"a", "b", "c", "d", "e"}
	fmt.Println("Before:", data)

	data = removeAtIndex(data, 2)

	fmt.Println("After: ", data)
}
