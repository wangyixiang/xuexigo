package btbasics5

/*
Q18. (1) Pointer arithmetic
1. In the main text on page 150 there is the following text:
…there is no pointer arithmetic, so if you write: *p++, it is
interpreted as (*p)++: first dereference and then increment
the value.

When you increment a value like this, for which types will it work?

2. Why doesn’t it work for all types?
*/

import "fmt"

func Q18_1() {
	fmt.Println(`It should work on 
	1) uint8 uint uint16 uint32 uint64
	2) int8 int int16 int32 int64
	3) float32 float32
	4) byte 
	`)
}

func Q18_2() {
	fmt.Println(`It doesn't work on 
	1) string
	2) bool
	3) self defined type struct
	4) type which not defined from numerical type
	5) error
	6) complex64, complex128
	`)
	fmt.Println("because not all type support ++ operater.")
}
