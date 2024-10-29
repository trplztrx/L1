package main

import (
	"fmt"
)

func intersection(set1, set2 []int) []int {
	var intersection []int

	setMap := make(map[int]struct{})
	for _, num := range set1 {
		setMap[num] = struct{}{}
	}

	for _, num := range set2 {
		if _, exists := setMap[num]; exists {
			intersection = append(intersection, num)
			delete(setMap, num)
		}
	}

	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	res := intersection(set1, set2)
	
	fmt.Println("Множество 1:", set1)
	fmt.Println("Множество 2:", set2)
	fmt.Println("Пересечение:", res)
}
