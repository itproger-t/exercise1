package main

import "fmt"

func binarySearch(arr []int, val int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (right + left) / 2
		if arr[mid] > val {
			right = mid - 1
		} else if arr[mid] < val {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 4, 7, 8, 11, 15}
	val := 7
	fmt.Println(binarySearch(arr, val))
	fmt.Println(binarySearch(arr, 1))
	fmt.Println(binarySearch(arr, 0))
	fmt.Println(binarySearch(arr, 11))
	fmt.Println(binarySearch(arr, 15))
	fmt.Println(binarySearch(arr, 16))
}
