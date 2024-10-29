package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var ttl int
	fmt.Printf("Enter time to live of program (sec): ")
	fmt.Scanln(&ttl)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * time.Duration(ttl))
	defer cancel()

	tasks := make(chan any)

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func(ctx context.Context, wg *sync.WaitGroup, tasks chan<- any) {
		defer wg.Done()

		ticker := time.NewTicker(time.Millisecond * 500)
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopping sender by timeout...")
				return
			case <-ticker.C:
				tasks <- rand.Intn(100)
			}
		}
	}(ctx, wg, tasks)

	go func(ctx context.Context, wg *sync.WaitGroup, tasks <-chan any) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Stopping receiver by timeout...")
				return
			case task := <-tasks:
				fmt.Println("Received value: ", task)
			}
		}
	}(ctx, wg, tasks)
	

	wg.Wait()
	close(tasks)
	fmt.Println("All goroutines have finished")
}