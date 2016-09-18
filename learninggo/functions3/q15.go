package functions3

/*
Q15. (1) Functions that return functions

1. Write a function that returns a function that performs a +2 on integers.
Name the function plusTwo. You should then be able do the following:
	p := plusTwo()
	fmt.Printf("%v\n", p(2))
Which should print 4. See section Callbacks on page 94 for information
about this topic.

2. Generalize the function from 1, and create a plusX(x) which returns
functions that add x to an integer.

*/
import "fmt"

func Q15_1() {
	af := q15_1(5)
	fmt.Printf("%d\n", af())

}

func q15_1(i int) func() int {
	return func() int {
		return i + 2
	}
}
func Q15_2() {
	af := q15_2(5)
	fmt.Printf("%d\n", af())
}

func q15_2(i int) func() int {
	return func() int {
		return i * 2
	}
}
