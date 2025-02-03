package main

import "fmt"

func main(){
	fmt.Println(add(1,2))
	fmt.Println(add(1,2,5))
	fmt.Println(add(1,25,1,4))
	handle(1, "abc")
    handle("abc", "xyz", 3)
}

func add(numbers ...int)int{
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}
func handle(params ...interface{}){
	fmt.Println("Handle func called with :")
	for _, it := range params{
		fmt.Printf("%v \n",it)
	}

}