package functions3

/*
Q13. (0) Minimum and maximum
1. Write a function that finds the maximum value in an int slice ([]int).
2. Write a function that finds the minimum value in an int slice ([]int).
*/

import "fmt"

func Q13_1() {
	as := []int{-1, 3, 8, 100, 1333, 55}
	fmt.Printf("%v\n", as)
	fmt.Printf("Max value is %d \n", q13_1(as))
	fmt.Printf("Min value is %d \n", q13_2(as))
}
func q13_1(s []int) int {
	if len(s) == 1 {
		return s[0]
	}
	maxi := s[0]
	for _, value := range s[1:] {
		if value > maxi {
			maxi = value
		}
	}

	return maxi
}

func q13_2(s []int) int {
	if len(s) == 1 {
		return s[0]
	}
	mini := s[0]
	for _, value := range s[1:] {
		if value < mini {
			mini = value
		}
	}

	return mini
}
