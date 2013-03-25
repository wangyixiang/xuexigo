// q3.go
package basics2

/*
Q3. (0) FizzBuzz
1. Solve this problem, called the Fizz-Buzz [30] problem:
Write a program that prints the numbers from 1 to 100. But
for multiples of three print “Fizz” instead of the number and
for the multiples of five print “Buzz”. For numbers which are
multiples of both three and five print “FizzBuzz”.
*/

import (
	"fmt"
)

const (
	F   = "Fizz"
	FI  = 3
	B   = "Buzz"
	BI  = 5
	FB  = "FizzBuzz"
	FBI = FI * BI
)

func Q3() {
	for i := 1; i <= 100; i++ {
		if i%FBI == 0 {
			fmt.Println(i, " ", FB)
			continue
		}
		if i%FI == 0 {
			fmt.Println(i, " ", F)
			continue
		}
		if i%BI == 0 {
			fmt.Println(i, " ", B)
			continue
		}
	}
}
