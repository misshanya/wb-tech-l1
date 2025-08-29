package main

import (
	"fmt"
	"math/big"
)

func main() {
	first := big.NewInt(1 << 30)
	second := big.NewInt(1 << 23)

	first.Mul(first, second)
	fmt.Println("Result of multiplication:", first)

	first.Div(first, second)
	fmt.Println("Result of divide:", first)

	first.Add(first, second)
	fmt.Println("Result of addition:", first)

	first.Sub(first, second)
	fmt.Println("Result of subtraction:", first)
}
