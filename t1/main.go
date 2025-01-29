package main

import "fmt"

func main() {
//    twoDSlice := make([][]int, 4)
////    twoDSlice[0] = []int{1, 2, 3}
////    twoDSlice[1] = []int{4, 5}
////	twoDSlice[2] = []int{1,5,7}
////	twoDSlice[3] = []int{2,3,4,5}
////
////
////  
////    fmt.Printf("Number of rows in slice: %d\n", len(twoDSlice))
////    fmt.Printf("Len of first row: %d\n", len(twoDSlice[0]))
////    fmt.Printf("Len of 3rd row: %d\n", len(twoDSlice[3]))
////    fmt.Println("Traversing slice")
////    for _, row := range twoDSlice {
////        for _, val := range row {
////            fmt.Print(val)
////        }
////		fmt.Println(" ")
  //  }
  letters := []string{"a", "b", "c", "d", "e"}
  for _ , let := range letters{
	fmt.Println(let)
  }
  i := 0
  for i = range letters{
	fmt.Println(letters[i])
	
  }
  source := []string{"san", "man", "tan"}
    result := find(source, "san")
    fmt.Printf("Item san found: %t\n", result)
    result = find(source, "can")
    fmt.Printf("Item san found: %t\n", result)
}
func find(source []string, target string) bool {
	for _, item := range source{
		if item == target {
			return true
		}
	}
	return false
}