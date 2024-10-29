package main

import "fmt"

func swap(arr []rune, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}

func reverse(str string) string {
	runes := []rune(str)

	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		swap(runes, i, j)
	}

	return string(runes)
}

func main() {
	strIn := "snow dog suns"
	fmt.Println("In:  ", strIn)

	strOut := reverse(strIn)
	fmt.Println("Out: ", strOut)
}