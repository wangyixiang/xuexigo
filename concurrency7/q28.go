package concurrency7

/*
Q28. (2) Fibonacci II
1. This is the same exercise as the one given page 101 in exercise 11. For
completeness the complete question:
The Fibonacci sequence starts as follows:
1; 1; 2; 3; 5; 8; 13; : : : Or in mathematical terms:
x 1 = 1; x 2 = 1; x n = x n − 1 + x n − 2 ∀ n > 2.
Write a function that takes an int value and gives that many
terms of the Fibonacci sequence.
But now the twist: You must use channels.
*/
import "fmt"

var ch2 chan float64

func Q28_1() {
	var i float64 = 1
	ch2 = make(chan float64)
	go q28_1(100)
	for i != 0 {
		i = <-ch2
		fmt.Printf("%d ", i)
	}
}

func q28_1(x float64) {
	a := 0.0
	b := 1.0
	for i := 1.0; i <= x; i++ {
		a, b = b, a+b
		ch2 <- b
	}
	fmt.Println("finished!")
	close(ch2)
}
