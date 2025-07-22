package main

import (
	"fmt"
	"strconv"
)

func oneZero(num, bit int) int {
	mask := 1 << bit
	return num &^ mask
}

func zeroOne(num, bit int) int {
	mask := 1 << bit
	return num ^ mask
}

func main() {
	bit, num := 6, 567
	num2 := zeroOne(num, bit)
	fmt.Println(num, num2, oneZero(num2, bit))

	fmt.Println(strconv.FormatInt(int64(num), 2))
	fmt.Println(strconv.FormatInt(int64(num^(1<<bit)), 2))
	fmt.Println(strconv.FormatInt(int64(num2&^(1<<bit)), 2))
}
