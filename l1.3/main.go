package main

import (
	"fmt"
	"sync"
)

func Workers(n int, wg *sync.WaitGroup, ch <-chan int) {
	for i := 1; i <= n; i++ {
		go func(j int) {
			defer wg.Done()
			for num := range ch {
				fmt.Printf("Worker: %d, value: %d\n", j, num)
			}
		}(i)
	}
}

func main() {
	var n int = 5
	wg := sync.WaitGroup{}
	ch := make(chan int, n)
	wg.Add(n)
	Workers(n, &wg, ch)
	for _, num := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15} {
		ch <- num
	}
	close(ch)
	wg.Wait()
}
