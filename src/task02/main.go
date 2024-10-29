package main

import (
	"fmt"
	"sync"
)

func main() {
	inputSlice := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	
	for _, num := range inputSlice {
		wg.Add(1)
		// num := num
		go func(num int) {
			defer wg.Done()
			fmt.Println(num * num)
		}(num)
	}
	wg.Wait()
}