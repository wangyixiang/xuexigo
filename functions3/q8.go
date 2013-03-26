package functions3

/*
Q8. (1) Scope
1. What is wrong with the following program?
1 package main
3 import "fmt"func main() {
6 for i := 0; i < 10; i++ {
7 fmt.Printf("%v\n", i)
8 }
9 fmt.Printf("%v\n", i)
10 }
*/

import "fmt"

func Q8_1() {
	fmt.Println("The \"i\" in line9 will not be vaild, because i only is valid in the for loop statements above.")
}
