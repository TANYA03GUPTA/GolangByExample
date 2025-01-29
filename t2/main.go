//maps
package main

import "fmt"

type animal interface {
	breathe()
	walk()
}
type mouse struct{
	age int32
}
type loin struct{
	age int32
}
func (m mouse) breathe() {
	fmt.Println("mouse breathe")
}
func (l loin) breathe(){
	fmt.Println("loin breathe")
}
func (l loin) walk(){
	fmt.Println("lion walks")
}
func(m mouse) walk(){
	fmt.Println("mouse walks")
}

func main(){
	//declare a map
	
	mp :=  map[int]string{
		1:"hello",
		4:"pick",
		5:"yellow",
	}
	for _, v:= range mp{
		fmt.Println(v)
	}
	var a animal
	a = mouse{age: 20}
	a.breathe()
    a.walk()
	var b animal
	b = loin{age: 12}
	b.breathe()
	b.walk()


}