package main

import (
	"fmt"
)

func reverseString(str string) string {
	reverse := []rune(str)
	ln := len(reverse)
	for i := 0; i < ln/2; i++ {
		reverse[i], reverse[ln-1-i] = reverse[ln-1-i], reverse[i]
	}
	return string(reverse)
}

func main() {
	str := "Lux©It"
	fmt.Println(str)
	fmt.Println(reverseString(str))
}
