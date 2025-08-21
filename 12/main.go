package main

import "fmt"

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}

	wordsMap := make(map[string]bool)
	for _, word := range words {
		wordsMap[word] = true
	}

	wordsUnique := make([]string, 0, len(wordsMap))
	for word := range wordsMap {
		wordsUnique = append(wordsUnique, word)
	}

	fmt.Println(wordsUnique)
}
