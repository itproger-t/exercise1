package main

import (
	"fmt"
	"strings"
)

func unique(str string) bool {
	set := make(map[string]struct{})
	for _, ch := range str {
		s := strings.ToLower(string(ch))
		if _, ok := set[s]; ok {
			return false
		}
		set[s] = struct{}{}
	}
	return true
}
func main() {
	fmt.Println(unique("abcd"))
	fmt.Println(unique("abCdefAaf"))
	fmt.Println(unique("aabcd"))
}
