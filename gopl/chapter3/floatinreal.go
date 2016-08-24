package main

import "fmt"

func main() {
	fmt.Printf("%v\n", float32(1 << 24 - 1))
	fmt.Printf("%v\n", float32(1 << 24))
	fmt.Printf("%v\n", float32(1 << 24 + 1))
	fmt.Printf("%v\n", float32(1 << 24 + 2))
}
