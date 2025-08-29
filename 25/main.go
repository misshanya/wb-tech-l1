package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("We are going to sleep for 2 seconds")
	sleep(2 * time.Second)
	fmt.Println("We just woke up")
}

func sleep(d time.Duration) {
	<-time.After(d)
}
