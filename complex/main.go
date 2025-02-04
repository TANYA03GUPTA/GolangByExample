package main
import (
    "fmt"
    "reflect"
    "unsafe"
)
func main() {
    var a float32 = 3
    var b float32 = 5
    
    //Initialize-1
    c := complex(a, b)
    
    //Initialize-2
    var d complex64
    d = 4 + 5i
    
    //Print Size
    fmt.Printf("c's size is %d bytes\n", unsafe.Sizeof(c))
    fmt.Printf("d's size is %d bytes\n", unsafe.Sizeof(d))
    
    //Print type
    fmt.Printf("c's type is %s\n", reflect.TypeOf(c))
    fmt.Printf("d's type is %s\n", reflect.TypeOf(d))
    
    //Operations on complex number
    fmt.Println(c+d, c-d, c*d, c/d)
}