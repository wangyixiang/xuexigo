// q7.go
package functions3

/*
Q7. (0) Integer ordering
1. Write a function that returns its (two) parameters in the right, numer-
ical (ascending) order:
f(7,2) → 2,7
f(2,7) → 2,7
*/

func Q7_1(x, y int) (nx, ny int) {
	if x > y {
		return x, y
	}
	return y, x
}
