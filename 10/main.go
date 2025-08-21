package main

import "fmt"

func main() {
	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := map[int][]float64{}

	for _, num := range nums {
		group := 10 * int(num/10)
		groups[group] = append(groups[group], num)
	}

	fmt.Println(groups)
}
