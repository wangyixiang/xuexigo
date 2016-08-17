package main

import (
	"flag"
	"fmt"
)

func main() {
	var num = flag.Int("num", 101, "upper bound of generator")
	flag.Parse()
	chYield := make(chan int)
	chSquare := make(chan int)

	go func() {
		for i := 1; i <= *num; i++ {
			chYield <- i
		}
		close(chYield)
	}()

	go func() {
		for {
			num, ok := <-chYield
			if !ok {
				close(chSquare)
				break
			}
			chSquare <- num * num
		}
	}()

	func() {
		for num := range chSquare {
			fmt.Print(num, "\t")
		}
		fmt.Println()
	}()

}
