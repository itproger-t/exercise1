package main

import "fmt"

func main() {
	ans := make(map[int64][]float64)
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 29.99, 30}
	for _, num := range arr {
		key := int64(num)
		key = (key - key%10)
		ans[key] = append(ans[key], num)
	}
	fmt.Println(ans)
}
