package main

import "fmt"

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	pivot := arr[0]
	var low, high []int
	for i := 1; i < len(arr); i++ {
		num := arr[i]
		if num < pivot {
			low = append(low, num)
		} else if num >= pivot {
			high = append(high, num)
		}
	}
	low = append(quickSort(low), pivot)
	low = append(low, quickSort(high)...)
	return low
}

func main() {
	arr := []int{4, 2, 1, 11, 5, 100, 56, 1, 1, 1, 111}
	fmt.Println(quickSort(arr))
}
