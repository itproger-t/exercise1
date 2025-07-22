package main

import (
	"fmt"
	"slices"
	"time"
)

func CompetitiveSquaring(arr []int) {
	ch := make(chan struct{}, 5)
	for i := range arr {
		ch <- struct{}{}
		go func(j int) {
			arr[j] *= arr[j]
			<-ch
		}(i)
	}
	time.Sleep(time.Duration(1) * time.Second)
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	CompetitiveSquaring(arr1)
	fmt.Println(slices.Equal(arr1, []int{1, 4, 9, 16, 25}), arr1)

	arr2 := []int{6, 7, 8, 9, 10}
	CompetitiveSquaring(arr2)
	fmt.Println(slices.Equal(arr2, []int{36, 49, 64, 81, 100}), arr2)
}
