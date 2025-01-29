package main

import (
	"fmt"
	_ "time"
)

func main(){
	//chan : bidirectional
	//chan <- only write 
	//<- chan only read
	ch := make(chan int, 3)
	fmt.Printf("Length of chan is : %d\n", len(ch))
	send(ch,4)
	//fmt.Println(<-ch)
	pr(ch)
	send(ch,16)
	fmt.Printf("Length of chan is : %d\n", len(ch))
    pr(ch)
	send(ch, 52)
	pr(ch)
	fmt.Printf("Length of chan is : %d\n", len(ch))
	send(ch,58)
	fmt.Printf("%d\n",cap(ch))
	ch2 := make(chan int)
	
	go func(){
		ch2 <- 43
		ch2 <- 52
	}()
	fmt.Println(<- ch2)

	ch3 := make(chan int, 1)
	ch3 <- 21
	val,ok := <-ch3

	fmt.Printf("Value is : %d\n ok : %t\n", val,ok)
	close(ch3)
	val,ok = <-ch3

	fmt.Printf("Value is : %d\n ok : %t\n", val,ok)
	
}
func process(ch chan<- int){
	ch <- 2
	ch <- 5
	ch <- 6
}
func pr(ch <-chan int){
	fmt.Println(<-ch)
}
func send(ch chan int,x int){
	fmt.Println("sending value to chan start")
	ch <- x
	fmt.Println("sending value finished")
}