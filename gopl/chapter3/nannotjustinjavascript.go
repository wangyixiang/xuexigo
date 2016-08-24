package main

import (
	"fmt"
	"math"
)

func main() {
	var z float64
	fmt.Printf("%v, %v, %v, %v, %v, %v\n", z, -z, 1/z, -1/z, z/z, z*z)
	fmt.Printf("%v, %v\n", math.NaN(), math.IsNaN(math.NaN()))
}
