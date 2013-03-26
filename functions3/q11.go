package functions3

/*
Q11. (1) Fibonacci
1. The Fibonacci sequence starts as follows: 1; 1; 2; 3; 5; 8; 13; : : : Or in
mathematical terms: x 1 = 1; x 2 = 1; x n = x n − 1 + x n − 2 ∀ n > 2.
Write a function that takes an int value and gives that many terms of
the Fibonacci sequence.
*/
import "fmt"

func Q11_1(value uint) {
	a := uint(1)
	b := uint(1)
	for i := uint(3); i <= value; i++ {
		fmt.Printf("%d %d\n", i, a+b)
		a, b = b, a+b
	}
}
