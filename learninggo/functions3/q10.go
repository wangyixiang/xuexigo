package functions3

/*
Q10. (1) Var args
1. Write a function that takes a variable number of ints and prints each
integer on a separate line

*/
import "fmt"

func Q10_1(a ...int) {
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}
