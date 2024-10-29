package main

import (
	"fmt"
	"strings"
)

func swap(arr []string, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func reverse(str string) string{
	words := strings.Fields(str)
	for i, j := 0, len(words) - 1; i < j; i, j = i + 1, j - 1 {
		swap(words, i, j)
	}

	return strings.Join(words, " ")
}

func main() {
	strIn := "snow dog sun"
	fmt.Println("In:  ", strIn)

	strOut := reverse(strIn)
	fmt.Println("Out: ", strOut)
}