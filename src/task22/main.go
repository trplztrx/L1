package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(2020399492019545)
	b := big.NewInt(1929492492949454)

	sum := new(big.Int)
	sub := new(big.Int)
	mul := new(big.Int)
	div := new(big.Float)

	sum.Add(a, b)
	sub.Sub(a, b)
	mul.Mul(a, b)

	aFloat := new(big.Float).SetInt(a)
	bFloat := new(big.Float).SetInt(b)
	div.Quo(aFloat, bFloat)

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("a + b: ", sum)
	fmt.Println("a - b: ", sub)
	fmt.Println("a * b: ", mul)
	fmt.Println("a / b: ", div)
}