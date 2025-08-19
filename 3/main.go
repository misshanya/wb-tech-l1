package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	amountOfWorkers := flag.Int("amount", 10, "Amount of workers to spawn")
	flag.Parse()

	ch := make(chan interface{})

	workerGroup := NewWorkerGroup(*amountOfWorkers, ch)
	workerGroup.Start()

	for {
		ch <- generateRandomString(10)
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}
}

func generateRandomString(length int) string {
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

type WorkerGroup struct {
	amountOfWorkers int
	ch              chan interface{}
}

func NewWorkerGroup(amountOfWorkers int, ch chan interface{}) *WorkerGroup {
	return &WorkerGroup{amountOfWorkers, ch}
}

func (g *WorkerGroup) Start() {
	for i := 0; i < g.amountOfWorkers; i++ {
		go func() {
			fmt.Println("Starting worker", i)
			for val := range g.ch {
				fmt.Println(val)
			}
		}()
	}
}
