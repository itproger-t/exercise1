package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func worker1(stop *bool) {
	for {
		if *stop {
			fmt.Println("Worker stopped by flag")
			return
		}
		fmt.Println("Working...")
		time.Sleep(500 * time.Millisecond)
	}
}

func startWorker1() {
	stop := false
	go worker1(&stop)

	time.Sleep(2 * time.Second)
	stop = true

	time.Sleep(1 * time.Second)
}

func worker2(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Worker stopped via channel")
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func startWorker2() {
	stopChan := make(chan struct{})
	go worker2(stopChan)

	time.Sleep(2 * time.Second)
	close(stopChan)

	time.Sleep(1 * time.Second)
}

func worker3(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped via context:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func startWorker3() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker3(ctx)

	time.Sleep(2 * time.Second)
	cancel()

	time.Sleep(1 * time.Second)
}

func worker4() {
	defer fmt.Println("Worker exited (via Goexit)")
	fmt.Println("Working...")
	time.Sleep(1 * time.Second)
	runtime.Goexit() // немедленно завершает выполнение текущей горутины
	fmt.Println("This will not be printed")
}

func startWorker4() {
	go worker4()

	time.Sleep(2 * time.Second)
}

func worker5() {
	for i := 0; i < 3; i++ {
		fmt.Println("Working...", i)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Worker finished via return")
	return
}

func startWorker5() {
	go worker5()

	time.Sleep(2 * time.Second)
}

func worker6() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Worker stopped with panic:", r)
		}
	}()

	for i := 0; i < 5; i++ {
		if i == 2 {
			panic("manual stop")
		}
		fmt.Println("Working...", i)
		time.Sleep(300 * time.Millisecond)
	}
}

func startWorker6() {
	go worker6()

	time.Sleep(2 * time.Second)
}

func main() {
	startWorker1()
	startWorker2()
	startWorker3()
	startWorker4()
	startWorker5()
	startWorker6()
}
