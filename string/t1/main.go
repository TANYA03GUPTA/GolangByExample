// atoi
//replaceall
//multiline
//compare
//substring chk
//get all words of a sentence
package main

import (
	"fmt"
	"strconv"
	"strings"
)
func main() {
    x := "1234"
	y1 := x
    val, err := strconv.Atoi(x)
	fmt.Println(y1)
    if err != nil {
        fmt.Printf("Supplied value %s is not a number\n", x)
    } else {
        fmt.Printf("Supplied value %s is a number with value %d\n", x, val)
    }

    y := "123b"
    val, err = strconv.Atoi(y)
    if err != nil {
        fmt.Printf("Supplied value %s is not a number\n", y)
    } else {
        fmt.Printf("Supplied value %s is a number with value %d\n", y, val)
    }
	st1 := "This is a sbuetifull garden"

	nospace := strings.ReplaceAll(st1," ","")
	fmt.Println(nospace)

	st2 := `this is myt 
	house
	where only
	Raju lievs`
	fmt.Println(st2)
	res := strings.Compare("tanya","Tanya")
	fmt.Println(res)
    prs := strings.Contains("Tanya", "xyz")
	fmt.Println(prs)
	rews := strings.Split("ab$cd$ef", "$")
    fmt.Println(rews)

	res2 := strings.Fields("af hello aap kayn ho ?")
    fmt.Println(res2)
}

