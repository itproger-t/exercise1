package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func worker(ctx context.Context, ch <-chan int) {
	for {
		select {
		case val := <-ch:
			fmt.Printf("Read: %d\n", val)
		case <-ctx.Done():
			log.Println("The program is complete")
			return
		}
	}
}

func main() {
	val := 0
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	go worker(ctx, ch)
	for {
		select {
		case ch <- val:
			val += 1
			time.Sleep(time.Duration(1) * time.Second)
		case <-ctx.Done():
			return
		}
	}
}
