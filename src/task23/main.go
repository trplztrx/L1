package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	n := len(slice)
	i := 2

	for j := i; j < n - 1; j++ {
		slice[j] = slice[j + 1]
	}
	slice = slice[:len(slice) - 1]
	fmt.Println(slice)
}