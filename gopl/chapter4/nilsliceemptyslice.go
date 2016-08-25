package main

import "fmt"

func main() {
	var anils []int
	anils = nil
	// it should panic
	fmt.Printf("%v\t%v\n", len(anils), cap(anils))
	fmt.Printf("%v\n", anils == nil)
	// panic
	// anils[0] = 1
	aemptys := []int{}
	fmt.Printf("%v\t%v\n", len(aemptys), cap(aemptys))
	fmt.Printf("%v\n", aemptys == nil)
	// panic
	// aemptys[0] = 1
}
