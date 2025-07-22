package main

import "fmt"

func main() {
	arr1 := map[int]struct{}{1: struct{}{}, 2: struct{}{}, 3: struct{}{}}
	arr2 := map[int]struct{}{2: struct{}{}, 3: struct{}{}, 4: struct{}{}}

	var result []int
	for key := range arr1 {
		if _, ok := arr2[key]; ok {
			result = append(result, key)
		}
	}
	fmt.Println(result)
}
