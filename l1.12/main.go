package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	var result []string
	for _, key := range arr {
		if _, ok := set[key]; !ok {
			result = append(result, key)
		}
		set[key] = struct{}{}
	}
	fmt.Println(result)
}
