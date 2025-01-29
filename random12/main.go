 package main
 
import (
    "fmt"
    "math/rand"
    "strings"
    "time"
)
var (
    lowerCharSet   = "abcdedfghijklmnopqrst"
    upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    specialCharSet = "!@#$%&*"
    numberSet      = "0123456789"
    allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)
func main(){
	rand.Seed(time.Now().Unix())
	minSplChar := 1
	minNum := 1
	minUpperCase :=11
	pswdLen := 8
	pswd := generatePswd(pswdLen, minSplChar,minNum, minUpperCase)
	fmt.Println(pswd)

	minSplChar = 2
	minNum = 2
	minUpperCase = 2
	pswdLen = 20
	pswd = generatePswd(pswdLen, minSplChar,minNum, minUpperCase)
	fmt.Println(pswd)

}


func generatePswd(pswdlen, minsplchar,minNum, minUpperCase int) string{
	var pswd strings.Builder
	for i:=0 ;i<minsplchar;i++{
		random := rand.Intn(len(numberSet))
        pswd.WriteString(string(numberSet[random]))
	}
	for i:=0 ;i<minNum;i++{
		rand1 := rand.Intn(len(specialCharSet))
		pswd.WriteString(string(specialCharSet[rand1]))
	}
	for i := 0; i < minUpperCase; i++ {
        random := rand.Intn(len(upperCharSet))
        pswd.WriteString(string(upperCharSet[random]))
    }

    remainingLength := pswdlen - minsplchar - minNum - minUpperCase
    for i := 0; i < remainingLength; i++ {
        random := rand.Intn(len(allCharSet))
        pswd.WriteString(string(allCharSet[random]))
    }
    inRune := []rune(pswd.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)

}