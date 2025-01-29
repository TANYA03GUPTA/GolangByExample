package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	go start()
	fmt.Println("hello star")
	
	fmt.Println("ednign")
	for i := 0;i<5;i++{
		go exec2(i)
	}
	time.Sleep(time.Second * 8)
	fmt.Println("ednign")
	///
	var wg sync.WaitGroup
	wg.Add(3)

	go sleep(&wg, time.Second*1)
	go sleep(&wg, time.Second*2)
	go sleep(&wg, time.Second*3)

	wg.Wait()
	fmt.Println("Alll gorout finished")

	ch4 := make(chan int,1)
	go sum(15,25,ch4)
	val := <- ch4
	fmt.Println(val)
	close(ch4)
}
func sum(x ,y int , res chan int){
	sm := x+y
	time.Sleep(time.Second*2)
	res <- sm 
	return 
}
func sleep(wg *sync.WaitGroup, t time.Duration){
	defer wg.Done()
	time.Sleep(t)
	fmt.Println("finish exec ")
}
func exec2(id int){
     fmt.Printf("id : %d \n",id)
}
func start(){
	fmt.Println("In gorout")
	//go st2()
}
func st2(){
	fmt.Println("In st2")
}