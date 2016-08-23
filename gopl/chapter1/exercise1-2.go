package main

import (
	"os"
	"fmt"
)

func main() {
	for i, arg := range os.Args[1:]{
		fmt.Printf("%d:\t%s\n", i, arg)
	}
}