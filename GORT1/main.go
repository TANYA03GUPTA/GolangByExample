package main

import (
    "fmt"
    "time"
)

func printMessage(msg string) {
    for i := 1; i <= 5; i++ {
        fmt.Println(msg, i)
        time.Sleep(time.Millisecond * 500)
    }
}
func sendNumbers(ch chan int) {
    for i := 1; i <= 5; i++ {
        ch <- i
    }
    close(ch)
}

func main() {
    go printMessage("Hello")
    go printMessage("World")

    time.Sleep(time.Second * 3)
	ch := make(chan int)

    go sendNumbers(ch)

    for num := range ch {
        fmt.Println("Received:", num)
    }
	time.Sleep(time.Second * 15)
    fmt.Println("Main function exits")
}
