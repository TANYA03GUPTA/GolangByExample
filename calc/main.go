package main

import (
	"fmt"
	"math"
)

type Calculator interface {
	Add(a, b float64) float64
	Subtract(a, b float64) float64
	Multiply(a, b float64) float64
	Divide(a, b float64) (float64, error)
	sin(a float64) float64
	cos(a float64) float64
	Sqrt(a float64) float64
}

type Normal struct{
	*scientific
}
func(n *Normal) Add(a,b float64)float64{
	return a+b
}
func (n *Normal)Subtract(a,b float64) float64{
	return a-b
}
func (n *Normal)Multiply(a,b float64) float64{
	return a*b
}
func(n *Normal) Divide(a,b float64)(float64, error){
	if b==0{
		//fmt.Errorf("cannot divide by zero")
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a/b, nil
}
type scientific struct{
	*Normal
}
func (sc *scientific)sin(a float64)float64{
	return math.Sin(a)
}

func (sc *scientific)cos(a float64)float64{
	return math.Cos(a)
}

func( sc * scientific) Sqrt(a float64) float64{
	return math.Sqrt(a)
}




func main(){
	var calc Calculator
	var choice int
	fmt.Println("Select calculator Operation:")
	fmt.Println("1. Normal Calculator")
	fmt.Println("2. Scientific Calculator")
	fmt.Print("Enter choice (1 or 2): ")
	fmt.Scan(&choice)

	if choice ==1 {
		calc = &Normal{}
	}else if choice == 2{
		calc = &scientific{}
	}else{ fmt.Println("invalid choice ")
	return}

	var a,b float64
	var op string
	fmt.Print(" enter 1st num :")
	fmt.Scan(&a)
	if choice == 1 {
		// Normal calculator
		fmt.Print("Enter operation (+, -, *, /): ")
		fmt.Scan(&op)
		fmt.Print("Enter second number: ")
		fmt.Scan(&b)

		switch op {
		case "+":
			fmt.Printf("Result: %.2f\n", calc.Add(a, b))
		case "-":
			fmt.Printf("Result: %.2f\n", calc.Subtract(a, b))
		case "*":
			fmt.Printf("Result: %.2f\n", calc.Multiply(a, b))
		case "/":
			result, err := calc.Divide(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %.2f\n", result)
			}
		default:
			fmt.Println("Invalid operation.")
		}
	}else if choice == 2{ 
		// Scientific calculator
		fmt.Print("Enter operation ( sin, cos, tan, log)")
		fmt.Scan(&op)
		switch op {  
			case "sin": 
			   fmt.Printf("sin() == %.2f \n" , a, calc.sin(a))
			case "cos":
				fmt.Printf("cos(%.2f) = %.2f\n", a, calc.cos(a))
			case "sqrt":
				fmt.Printf("log(%.2f) = %.2f\n", a, calc.Sqrt(a))
			default:
				fmt.Println("Invalid operation.")
		}

	}
}