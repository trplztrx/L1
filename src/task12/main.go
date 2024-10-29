package main

import "fmt"

func main() {
	seq := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]struct{})

	for _, str := range seq {
		set[str] = struct{}{}
	}
	
	fmt.Println("In:  ", seq)
	fmt.Println("Out: ", set)
}