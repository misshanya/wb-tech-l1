package main

import (
	"fmt"
	"unicode"
)

func main() {
	text := "abcd"
	fmt.Printf("%s: %v\n", text, areCharsUnique(text))

	text = "abCdefAaf"
	fmt.Printf("%s: %v\n", text, areCharsUnique(text))
}

func areCharsUnique(s string) bool {
	runes := []rune(s)

	chars := make(map[rune]bool)

	for _, r := range runes {
		char := unicode.ToLower(r)
		if chars[char] {
			return false
		}
		chars[char] = true
	}

	return true
}
