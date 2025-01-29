package main

import (
	"fmt"
	"time"
)

func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go goOne(ch1)
	go goTwo(ch2)
	

	select{
	case msg1 := <-ch1:
		fmt.Println("Received from ch1:", msg1)
	case msg2 := <-ch2:
		time.Sleep(time.Second*10)
		fmt.Println("recieved froim ch2",msg2)	
	}
}


func goOne(ch chan string){
	ch <- "from 1st gorout"
}

func goTwo(ch chan string){
	ch <- "from 2nd disney"
}