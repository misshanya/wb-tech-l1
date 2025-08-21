package main

import "fmt"

func main() {
	a := 10
	b := 5

	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Println("a is", a)
	fmt.Println("b is", b)
}
