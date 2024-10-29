package main

import (
	"fmt"
	"math/rand"
)

func genRandArr(size int) ([]int) {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(10) + 1
	}
	return  arr
}

func main() {
	size := 10

	arr := genRandArr(size)

	fmt.Println(arr)
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, num := range arr {
			ch1<- num
		}
		close(ch1)
	}()

	go func() {
		for num := range ch1 {
			ch2<- num * num
		}
		close(ch2)
	}()

	for num := range ch2 {
		fmt.Println(num)
	}
}