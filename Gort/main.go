package main 

import (
	"fmt"
	"time"
)

type reportcard struct{
	name string
	totalmarks int
	percentage float32
}
func main(){

	ch := make(chan bool)
	go test(ch)

	time.Sleep(time.Second * 1)
	fmt.Printf("send value : %t \n", true)
	ch <-true

	time.Sleep(time.Second*3)
    
	reschan := make(chan reportcard,2)
	go helper("Rohit Garg",465,reschan)
    go helper("RMohit Gupta",485,reschan)
	time.Sleep(time.Second*5)
	close(reschan)
	for res:= range reschan{
		//res := <-reschan
	fmt.Printf("Name of child is : %v \n",res.name)
	fmt.Printf("Total Marks of child is : %d \n",res.totalmarks)
	fmt.Printf("Percentage scored : %f \n",res.percentage)
	//i++
	}
	
    
}
func helper(nm string, ttl int, res chan reportcard){
	percent := ttl /5
	res1 := reportcard{name: nm, totalmarks: ttl, percentage: float32(percent)}
	res <- res1
	return
}
func test(ch chan bool) {
	var va1 = <- ch
	fmt.Printf("recieved value : %t \n",va1)
	fmt.Println("doing some work")
}