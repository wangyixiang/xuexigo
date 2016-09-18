package interfaces6

/*
 (2) Interfaces and max()
1. In exercise Q13 we created a max function that works on a slice of inte-
gers. The question now is to create a program that shows the maximum
number and that works for both integers and floats. Try to make your
program as generic as possible, although that is quite difficult in this
case.
*/

import "fmt"

/*
!!wang!! In this Question I meet a lot interesting situation.
*/
func Q26_1() {
	is := []int{1, 2, 3, 4, 5, 6}
	fs := []float32{1.1, 1.2, 1.3, 1.4, 1.5, 1.6}
	/*Because Go can not easily convert to a slice of interface,
	but just convert an interface is easy, so Go does not implicitly
	convert slices for you.
	*/
	//fmt.Printf("max of is %d\n", max(is))
	//fmt.Printf("max of fs %d\n", max(fs))
	//fmt.Printf("is %v", is)
	//fmt.Printf("fs %v", fs)
	//fmt.Println uses the reflect mechanism to get the value of is
	//fmt.Println(is)

	fmt.Println(max(is))
	fmt.Println(max(fs))
	fmt.Println(max(nil))
}

/*!!wang!! unlike python, in Go lang there's no an type for type, so reflect will be used heavily if needed.
 */
func max(s interface{}) interface{} {
	switch s.(type) {
	case []int:
		a := s.([]int)
		result := a[0]
		for i := 0; i < len(a); i++ {
			if result < a[i] {
				result = a[i]
			}

		}
		return result
	case []float32:
		a := s.([]float32)
		result := a[0]
		for i := 0; i < len(a); i++ {
			if result < a[i] {
				result = a[i]
			}

		}
		return result
	case []float64:
		a := s.([]float64)
		result := a[0]
		for i := 0; i < len(a); i++ {
			if result < a[i] {
				result = a[i]
			}

		}
		return result
	}
	return nil
}

/*
Problem one
func max(sl []interface{}) interface{} {
	mv := sl[0]
	for _, value := range sl[1:] {
		switch value.(type) {
		case int:
			switch mv.(type) {
			case int:
				if mv < value {
					mv = value
				}
			case float32:
				if mv < float32(value) {
					mv = value
				}
			}
		case float32:
			switch mv.(type) {
			case int:
				if float32(mv) < value {
					mv = value
				}
			case float32:
				if mv < value {
					mv = value
				}
			}
		}
	}
	return mv
}
*/
