package main

import "fmt"

func main() {
	word := "главрыба"
	reversedWord := reverseString(word)

	fmt.Printf("Reversed %s is %s\n", word, reversedWord)
}

func reverseString(s string) string {
	result := []rune(s)

	l := 0
	r := len(result) - 1

	for l < r {
		result[l], result[r] = result[r], result[l]
		l++
		r--
	}

	return string(result)
}
