package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	list := strings.Split(str, " ")
	n := len(list)
	for i := 0; i < n/2; i++ {
		list[i], list[n-1-i] = list[n-1-i], list[i]
	}
	return strings.Join(list, " ")
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
	fmt.Println(reverseWords("snow"))
}
