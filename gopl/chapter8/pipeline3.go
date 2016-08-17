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

	go generator(*num, chYield)

	go square(chYield, chSquare)

	func() {
		for num := range chSquare {
			fmt.Print(num, "\t")
		}
		fmt.Println()
	}()

}

func generator(num int, out chan int) {
	for i := 1; i <= num; i++ {
		((chan <- int)(out)) <- i
	}
	close(out)
}

func square(in <- chan int, out chan <- int) {
	// the unidirectional channel are not allow to convert back.
	// ((chan int)(in)) <- 200
	for {
		num, ok := <-in
		if !ok {
			close(out)
			break
		}
		out <- num * num
	}
}