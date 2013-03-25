// basics2 project q2.go
package basics2

/*

Q2. (0) For-loop
1. Create a simple loop with the for construct. Make it loop 10 times and
print out the loop counter with the fmt package.

2. Rewrite the loop from 1. to use goto. The keyword for may not be used.

3. Rewrite the loop again so that it fills an array and then prints that array
to the screen.

*/

import (
	"fmt"
)

func Q2_1() {
	fmt.Println("Q2_1()")
	for i := 1; i < 11; i += 1 {
		fmt.Println(i)
	}
}

func Q2_2() {
	fmt.Println("Q2_2()")
	i := 1
b1:
	fmt.Println(i)
	if i < 10 {
		i += 1
		goto b1
	}
}

func Q2_3() {
	fmt.Println("Q2_3()")
	var iarray [10]int
	i := 1
	iarray[i-1] = i
b1:
	fmt.Println(iarray)
	if i < 10 {
		i += 1
		iarray[i-1] = i
		goto b1
	}
}
