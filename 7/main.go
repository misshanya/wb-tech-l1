package main

import (
	"fmt"
	"sync"
)

func main() {
	concurrentMap := NewConcurrentMap[string, string]()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("setting somekey = somevalue from 1 goroutine")
		concurrentMap.Set("somekey", "somevalue")

		fmt.Println("setting anotherkey = anothervalue from 1 goroutine")
		concurrentMap.Set("anotherkey", "anothervalue")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("setting somekey = somesomeval from 2 goroutine")
		concurrentMap.Set("somekey", "somesomeval")

		fmt.Println("setting anotherkey = anothervalval from 2 goroutine")
		concurrentMap.Set("anotherkey", "anothervalval")
	}()

	wg.Wait()
}

type ConcurrentMap[K comparable, V any] struct {
	data map[K]V
	mu   sync.RWMutex
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		data: make(map[K]V),
	}
}

func (c *ConcurrentMap[K, V]) Set(k K, v V) {
	c.mu.Lock()
	c.data[k] = v
	c.mu.Unlock()
}

func (c *ConcurrentMap[K, V]) Get(k K) (V, bool) {
	c.mu.RLock()
	v, ok := c.data[k]
	c.mu.RUnlock()
	return v, ok
}
