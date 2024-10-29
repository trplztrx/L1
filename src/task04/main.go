package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Worker(id int, tasks <-chan interface{}, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d finished by signal\n", id)
			return
		case task := <-tasks:
			// time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			fmt.Printf("Worker %d processed value: ", id)
			fmt.Println(task)
		}
	}
}
func main() {
	var numWorkers int
	fmt.Printf("Enter number of workers: ")
	fmt.Scanln(&numWorkers)
	// numWorkers = 10
	tasks := make(chan interface{})

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go Worker(i, tasks, ctx, &wg)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		ticker := time.NewTicker(time.Millisecond * 500)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				close(tasks)
				return
			case <-ticker.C:
				tasks <- rand.Intn(100)
			}
		}
	}()
		
	<-sigs
	fmt.Println("Stopping...")
	cancel()
	
	wg.Wait()
	fmt.Println("All workers have finished")
}