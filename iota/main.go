package main

import "fmt"

const (
    a = iota
    b
    c
)
func main() {
    fmt.Println(a)
    fmt.Println(a)
    fmt.Println(c)
}