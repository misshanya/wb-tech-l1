package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup
	wg.Add(len(nums))

	for _, num := range nums {
		go func() {
			defer wg.Done()
			fmt.Println(num * num)
		}()
	}

	wg.Wait()
}
