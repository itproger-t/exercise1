package main

import (
	"fmt"
	"strings"
)

/*
Мы создаем переменную которая является подстрокой(100 элементов)
строки(1024 элемента), проблема этого кода состоит в том, что
обе строки используют один резервный массив, из-за чего может произойти
утечка памяти
var justString string

func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

var justString string

func createHugeString(size int) string {
	return strings.Repeat("A", size)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
	fmt.Println(justString)
}
