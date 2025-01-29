//shuffle a string
package main
import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().Unix())

    in := "abcdedf£"

    inRune := []rune(in)
    rand.Shuffle(len(inRune), func(i, j int) {
        inRune[i], inRune[j] = inRune[j], inRune[i]
    })
    fmt.Println(string(inRune))

    rand.Shuffle(len(inRune), func(i, j int) {
        inRune[i], inRune[j] = inRune[j], inRune[i]
    })
	ir := rand.Intn(len(in))
    fmt.Println(string(inRune))
	fmt.Printf("%d \n", in[ir])
	fmt.Println(rand.Perm(12))
	rangeLower := 5
    rangeUpper := 20
	randomNum := rangeLower + rand.Intn(rangeUpper-rangeLower+1)
    fmt.Println(randomNum)
}