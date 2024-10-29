package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker1(stopChan chan struct{}) {
	for {
		select {
		case <-stopChan:
			fmt.Println("Worker 1 stopped")
			return
		default:
			fmt.Println("Worker 1 is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func stopWorker1() {
	stopChan := make(chan struct{})
	go worker1(stopChan)
	time.Sleep(3 * time.Second)
	close(stopChan)
	time.Sleep(1 * time.Second)
}

func worker2(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker 2 stopped by context")
			return
		default:
			fmt.Println("Worker 2 is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func stopWorker2() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker2(ctx)
	time.Sleep(3 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

var stop3 bool

func worker3() {
	for {
		if stop3 {
			fmt.Println("Worker 3 stopped")
			return
		}
		fmt.Println("Worker 3 is working...")
		time.Sleep(1 * time.Second)
	}
}

func stopWorker3() {
	go worker3()
	time.Sleep(3 * time.Second)
	stop3 = true
	time.Sleep(1 * time.Second)
}

func worker4() {
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Println("Worker 4 stopped by timeout")
			return
		default:
			fmt.Println("Worker 4 is working...")
			time.Sleep(1 * time.Second)
		}
	}
}

func stopWorker4() {
	go worker4()
	time.Sleep(7 * time.Second)
}

var stop5 bool
var mu sync.Mutex

func worker5() {
	for {
		mu.Lock()
		if stop5 {
			mu.Unlock()
			fmt.Println("Worker 5 stopped")
			return
		}
		mu.Unlock()
		fmt.Println("Worker 5 is working...")
		time.Sleep(1 * time.Second)
	}
}

func stopWorker5() {
	go worker5()
	time.Sleep(3 * time.Second)
	mu.Lock()
	stop5 = true
	mu.Unlock()
	time.Sleep(1 * time.Second)
}

func worker6(resultChan chan error) {
	for i := 0; i < 3; i++ {
		fmt.Println("Worker 6 is working...")
		time.Sleep(1 * time.Second)
	}
	resultChan <- nil
}

func stopWorker6() {
	resultChan := make(chan error)
	go worker6(resultChan)
	err := <-resultChan
	if err == nil {
		fmt.Println("Worker 6 finished successfully")
	}
}

func main() {
	fmt.Println("Starting Worker 1 (Channel)")
	stopWorker1() // Канал

	fmt.Println("Starting Worker 2 (Context)")
	stopWorker2() // Контекст

	fmt.Println("Starting Worker 3 (Global Variable)")
	stopWorker3() // Глобальная переменная

	fmt.Println("Starting Worker 4 (Timeout)")
	stopWorker4() // Таймер

	fmt.Println("Starting Worker 5 (Mutex)")
	stopWorker5() // Мьютекс

	fmt.Println("Starting Worker 6 (Error Channel)")
	stopWorker6() // Канал ошибки
}
