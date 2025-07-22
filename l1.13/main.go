package main

import "fmt"

func main() {
	//№1
	a, b := 23, 34
	fmt.Println(a, b)

	a, b = b, a
	fmt.Println(a, b)

	//№2
	c, d := 45, 56
	fmt.Println(c, d)
	c = c + d
	d = c - d
	c = c - d
	fmt.Println(c, d)
}
