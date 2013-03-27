package btbasics5

/*
Q19. (2) Map function with interfaces
1. Use the answer from exercise Q12, but now make it generic using in-
terfaces. Make it at least work for ints and strings.
*/
import "fmt"

type e interface{}

func Q19_1() {
	ai := []e{1, 2, 3, 4, 5}
	fi := func(i e) e {
		switch i.(type) {
		case int:
			return i.(int) * 2
		}
		return i
	}
	fmt.Printf("%v\n", q19_1(ai, fi))
}

//!!wang!! type switch
func q19_1(a []e, f func(e) e) []e {
	result := make([]e, len(a))
	for i, value := range a {
		result[i] = f(value)
	}

	return result
}
