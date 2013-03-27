package functions3

/*
Q14. (1) Bubble sort
1. Write a function that performs a bubble sort on a slice of ints. From
[31]:
It works by repeatedly stepping through the list to be sorted,
comparing each pair of adjacent items and swapping them
if they are in the wrong order. The pass through the list
is repeated until no swaps are needed, which indicates that
the list is sorted. The algorithm gets its name from the way
smaller elements “bubble” to the top of the list.
[31] also gives an example in pseudo code:
procedure bubbleSort( A : list of sortable items )do
swapped = false
for each i in 1 to length(A) - 1 inclusive do:
if A[i-1] > A[i] then
swap( A[i-1], A[i] )
swapped = true
end if
end for
while swapped
end procedure
*/

import "fmt"

func Q14_1() {
	asi := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	newasi := q14_1(asi)
	fmt.Printf("asi %v\n", asi)
	fmt.Printf("newasi after q14_1 %v\n", newasi)
	q14_1_p(&asi)
	fmt.Printf("asi after q14_1_p %v\n", asi)
}

func q14_1(s []int) []int {
	slen := len(s)
	i := 1
	for ; i < slen-1; i++ {
		for j := 0; j < slen-i; j++ {
			(s)[j], (s)[j+1] = judge((s)[j], (s)[j+1])
			(s)[j], (s)[j+1] = judge((s)[j], (s)[j+1])
		}
	}

	return s
}

//!!wang!! If I use pointer of s as the argument, how that can be?
func q14_1_p(s *[]int) {
	slen := len(*s)
	i := 1
	for ; i < slen-1; i++ {
		for j := 0; j < slen-i; j++ {
			(*s)[j], (*s)[j+1] = judge((*s)[j], (*s)[j+1])
		}
	}

}

func judge(x, y int) (int, int) {
	if x > y {
		return y, x
	}
	return x, y
}
