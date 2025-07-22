package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: shutting down...\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(time.Duration(2) * time.Second)
		}
	}
}

func main() {
	/*
		Обоснование: создаем контекст, чтобы удобно завершить все горутины,
		в главной горутине создаем сигнальный канал и останавливаем главную
		горутину до тех пор, пока в сигнальный канал не придут данные,
		после отменяем контекс и ждем завершения всех горутин
	*/
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(ctx, i, &wg)
	}

	signChan := make(chan os.Signal, 1)
	signal.Notify(signChan, syscall.SIGINT, syscall.SIGTERM)

	<-signChan
	fmt.Println("\nReceived interrupt signal. Shutting down...")
	cancel()
	wg.Wait()
	fmt.Println("All workers stopped. Exiting.")
}
