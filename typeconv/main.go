//convert float32->float64

package main

import "fmt"

func main() {
    var a float64 = 12
    var b int = int(a)
    fmt.Printf("Underlying Type of b: %T\n", b)

    b2 := int(a)
    fmt.Printf("Underlying Type of b2: %T\n", b2)
	c2 := float64(b2)
	fmt.Printf("%T",c2)
	c1 := int(45)
    c23 := float32(c1)
	fmt.Printf("%T",c23)
}