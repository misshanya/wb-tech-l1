package main

import "fmt"

func main() {
	words := "snow dog sun"
	reversed := reverseWordsInString(words)

	fmt.Println("Reversed is", reversed)
}

func reverseWordsInString(s string) string {
	result := []rune(s)
	reverseString(result)

	wordStarts := 0
	for i, char := range result {
		if char == ' ' || i == len(result)-1 {
			l := wordStarts

			var r int
			if char == ' ' {
				r = i - 1
			} else {
				r = i
			}

			for l < r {
				result[l], result[r] = result[r], result[l]
				l++
				r--
			}
			wordStarts = i + 1
		}
	}

	return string(result)
}

func reverseString(s []rune) {
	l := 0
	r := len(s) - 1

	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
