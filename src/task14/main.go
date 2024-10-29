package main

import (
	"fmt"
)

func checkType(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	var a interface{}

	a = 42
	checkType(a)

	a = "kostya shishkin"
	checkType(a)

	a = true
	checkType(a)

	a = make(chan int)
	checkType(a)

	a = make(chan bool)
	checkType(a)
}
