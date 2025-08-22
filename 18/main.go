package main

import (
	"fmt"
	"sync"
)

func main() {
	c := counter{
		mu: &sync.Mutex{},
	}

	var wg sync.WaitGroup

	for range 20 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()

	fmt.Println("Count is", c.GetCount())
}

type counter struct {
	count int
	mu    *sync.Mutex
}

func (c *counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *counter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
