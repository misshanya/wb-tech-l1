package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4, 3, 3}

	intersectionMap := map[int]bool{}

	mapA := map[int]bool{}
	for _, num := range a {
		mapA[num] = true
	}

	for _, num := range b {
		if _, ok := mapA[num]; ok {
			intersectionMap[num] = true
		}
	}

	intersection := make([]int, 0, len(intersectionMap))

	for num := range intersectionMap {
		intersection = append(intersection, num)
	}

	fmt.Println(intersection)
}
