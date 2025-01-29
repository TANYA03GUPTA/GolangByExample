package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go goOne(ch1)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout")
	}
}

func goOne(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "From goOne goroutine"
}