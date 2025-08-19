package main

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Use context with timeout to stop program after configured time
	timeoutDuration := 10 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	ch := make(chan interface{})

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		writer(ctx, ch)
	}()

	go func() {
		defer wg.Done()
		reader(ch)
	}()

	wg.Wait()
}

// writer writes pseudo-data to the given channel
func writer(ctx context.Context, ch chan interface{}) {
	idx := 0
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- "some data " + strconv.Itoa(idx)
			idx++
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		}
	}
}

// reader reads data from the given channel
func reader(ch chan interface{}) {
	for val := range ch {
		fmt.Printf("Got data: %v\n", val)
	}
}
