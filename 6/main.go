package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// By condition
	wg.Add(1)
	smth := "smth"
	go func() {
		defer wg.Done()
		defer fmt.Println("goroutine 1 done")
		if smth == "smth" {
			return
		}
	}()

	// By channel
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		defer fmt.Println("goroutine 2 done")
		<-ch
	}()
	ch <- 1

	// By channel closing
	wg.Add(1)
	ch = make(chan int)
	go func() {
		defer wg.Done()
		defer fmt.Println("goroutine 3 done")
		for {
			if _, ok := <-ch; !ok {
				return
			}
		}
	}()
	close(ch)

	// By context
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer wg.Done()
		defer fmt.Println("goroutine 4 done")
		<-ctx.Done()
	}()
	cancel()

	// By runtime.Goexit()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("goroutine 5 done")
		runtime.Goexit()
	}()

	wg.Wait()
}
