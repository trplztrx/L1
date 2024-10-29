package main

import (
	"fmt"
	"sync"
)
func main() {
	myMap := make(map[any]any)

	var wg sync.WaitGroup
	mu := new(sync.Mutex)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			mu.Lock()
			myMap[i] = n * n
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	for key,val := range myMap {
		fmt.Println(key, val)
	}
}