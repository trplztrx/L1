package main

import "fmt"

func BinSearch(arr []int, target int) int {
	left := 0
	right := len(arr)

	for left < right {
		mid := left + (right - left) / 2
		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11}
	target := 10

	fmt.Println("In:  ", arr)
	fmt.Println("Out: ", BinSearch(arr, target))
	// fmt.Println("Out right: ", RBinSearch(arr, target))
}