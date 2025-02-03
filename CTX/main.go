package main

import (
	"context"
	"fmt"
	"time"
)

func main(){
   //declare context 
   ctx := context.Background()
   cancelctx,cancel = context.WithCancel(ctx)
   go task(cancelctx)
   time.Sleep(time.Second*3)
   cancel()
   time.Sleep(time.Second*1)

}
func task(ctx context.Context){
	i :=1
	for{
		select{
		case <-ctx.Done():
			fmt.Println("task cancelled")
			return
		default:
			fmt.Println(i)
			i++
			time.Sleep(time.Second*1)	
		}
	}
}

--------


func main(){
	ctx := context.Background()
	cancelctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	go task1(cancelctx)
	time.Sleep(time.Second*4)
}
func task1(ctx context.Context){
	i:=1
	for{
		select{
			case <-ctx.Done():
				fmt.Println("task1 cancelled")	
			default:
				fmt.Println(i)
				time.Sleep(time.Second*1)
				i++	
		}
	}
}
-----
package main

import "fmt"

type Person struct {
    Name string
    Age  int
    City string
}

func main() {
    p := Person{Name: "Alice", Age: 25, City: "New York"}
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
    fmt.Println("City:", p.City)
}