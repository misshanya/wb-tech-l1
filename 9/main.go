package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numsCh := make(chan int)
	doubleCh := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range nums {
			numsCh <- num
		}
		close(numsCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range numsCh {
			doubleCh <- num * 2
		}
		close(doubleCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range doubleCh {
			fmt.Println(num)
		}
	}()

	wg.Wait()
}
