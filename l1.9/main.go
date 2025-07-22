package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)
	go func(ch1 <-chan int, ch2 chan<- int) {
		for num := range ch1 {
			ch2 <- num * 2
		}
		defer close(ch2)
	}(ch1, ch2)

	go func(ch <-chan int) {
		for num := range ch {
			fmt.Println(num)
		}
	}(ch2)

	for _, num := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		ch1 <- num
	}
	close(ch1)
	time.Sleep(time.Duration(1) * time.Second)
}
