package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		_ = 1 << (10 * iota)
		KB
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)
	fmt.Printf("KB %T %v\n", KB, KB)
	fmt.Printf("MB %T %v\n", MB, MB)
	fmt.Printf("GB %T %v\n", GB, GB)
	fmt.Printf("TB %T %v\n", TB, TB)
	fmt.Printf("PB %T %v\n", PB, PB)
	fmt.Printf("EB %T %v\n", EB, EB)
	//fmt.Printf("ZB %T %v\n", ZB, ZB)
	//fmt.Printf("YB %T %v\n", YB, YB)
	//fmt.Printf("ZB %T %v\n", float64(ZB), float64(ZB))
	//fmt.Printf("YB %T %v\n", float64(YB), float64(YB))
	fmt.Printf("Pi %T %v\n", math.Pi, math.Pi)
}
