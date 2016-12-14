package main

import "fmt"

func sum(vars ...int) int {
	var result int
	for _, v := range vars {
		result += v
	}
	return result
}

func main() {
	intarray := [...]int{1,2,3,4,5,6,7,8,9}
	fmt.Printf("%T\n", intarray)
	fmt.Printf("%T\n", intarray[:])
	fmt.Println(sum(intarray[:]...))
}
