package main

import "fmt"

func main() {
	var f float64 = 212
	fmt.Printf("%T\t%v\n", (f - 32) * 5 / 9, (f - 32) * 5 / 9)
	fmt.Printf("%T\t%v\n", 5 / 9 * (f - 32), 5 / 9 * (f - 32))
	fmt.Printf("%T\t%v\n", 5.0 / 9.0 * (f - 32), 5.0 / 9.0 * (f - 32))
}
