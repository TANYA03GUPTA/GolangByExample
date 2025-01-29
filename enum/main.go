package main 

import "fmt"
type Size int

const (
	small = Size(iota)
	medium
	large
	Xlarge
)

func main(){
	var m Size =3
	m.toString()
}

func (s Size) toString(){
	switch s {
		case small:
			  fmt.Println("small")
	    case medium: 
		fmt.Println("medium")
		case large :
			fmt.Println("large")
		case Xlarge:
			fmt.Println("Xlarge")
		default:
			fmt.Println("unknown")
	}
}