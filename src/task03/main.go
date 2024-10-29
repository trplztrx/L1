package main

import (
	"fmt"
	"sync"
)
func main() {
	inputSlice := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	sum := 0
	mu := new(sync.Mutex)
	for _, num := range inputSlice {
		wg.Add(1)
		// num := num
		go func(num int) {
			defer wg.Done()
			sq := num * num
			mu.Lock()
			sum += sq
			mu.Unlock()
		}(num)
	}
	wg.Wait()
	fmt.Println(sum)
}