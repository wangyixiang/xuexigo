// q5.go
package basics2

/*
Q5. (1) Average

1. Write code to calculate the average of a float64 slice. In a later exercise
(Q6 you will make it into a function.

*/
import (
	"fmt"
)

func Q5_1() {

	var f64slice = make([]float64, 5)
	f64slice[0] = 1
	f64slice[1] = 2
	f64slice[2] = 3
	f64slice[3] = 4
	f64slice[4] = 5

	fmt.Printf("%f\n", q5_1(f64slice))

}

func q5_1(f64s []float64) (average float64) {
	for i := 0; i < len(f64s); i += 1 {
		average += f64s[i]
	}

	return average / float64(len(f64s))

}
