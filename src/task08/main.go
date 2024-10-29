package main

import (
	"fmt"
	// "math/rand"
)

func setBitToOne(num int64, i uint) (int64) {
	return num | (1 << i)
}

func setBitToZero(num int64, i uint) (int64) {
	return num &^ (1 << i)
}
func main() {
	var x int64
	var i uint
	fmt.Print("Число (int64): ")
	fmt.Scanln(&x)

	fmt.Print("Число i: ")
	fmt.Scanln(&i)

	fmt.Printf("До установки битов: %d, %b\n", x, x)

	x = setBitToOne(x, i)
	fmt.Printf("После установки в %d-го бита в 1: %d, %b\n", i, x, x)

	x = setBitToZero(x, i)
	fmt.Printf("После установки в %d-го бита в 0: %d, %b\n", i, x, x)
}