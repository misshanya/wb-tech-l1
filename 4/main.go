package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func main() {
	amountOfWorkers := flag.Int("amount", 10, "Amount of workers to spawn")
	flag.Parse()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	ch := make(chan interface{})

	workerGroup := NewWorkerGroup(*amountOfWorkers, ch)
	workerGroup.Start()

	// Main logic: send some data to workers
	go func() {
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				ch <- generateRandomString(10)
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			}
		}
	}()

	<-ctx.Done()
	fmt.Println("Shutting down")
	workerGroup.Stop()
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
	stopWg          *sync.WaitGroup
}

func NewWorkerGroup(amountOfWorkers int, ch chan interface{}) *WorkerGroup {
	return &WorkerGroup{
		amountOfWorkers: amountOfWorkers,
		ch:              ch,
		stopWg:          &sync.WaitGroup{},
	}
}

func (g *WorkerGroup) Start() {
	for i := 0; i < g.amountOfWorkers; i++ {
		g.stopWg.Add(1)
		go func() {
			defer g.stopWg.Done()
			fmt.Println("Starting worker", i)
			for val := range g.ch {
				fmt.Println(val)
			}
			fmt.Println("Finished worker", i)
		}()
	}
}

func (g *WorkerGroup) Stop() {
	g.stopWg.Wait()
}
