// functions project q6.go
package functions3

/*
Q6. (0) Average
1. Write a function that calculates the average of a float64 slice.
*/
import ()

func Q6_1(f64s []float64) (average float64) {
	for i := 0; i < len(f64s); i += 1 {
		average += f64s[i]
	}

	return average / float64(len(f64s))

}
