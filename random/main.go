package main

import (
	"fmt"
	"math/rand"
	"time"
)

//func Intn(n int)int

func main(){
	//fmt.Println(Intn(5))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))

	rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Float32())
	fmt.Println(rand.Uint32())
}