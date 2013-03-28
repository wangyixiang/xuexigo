package btbasics5

/*
Q20. (1) Pointers

1. Suppose we have defined the following structure:
type Person struct {
name string
age int
}
What is the difference between the following two lines?
var p1 Person
p2 := new(Person)

2. What is the difference between the following two allocations?
func Set(t *T) {
x = t
}
and
func Set(t T) {
x= &t
}

*/

import "fmt"

func Q20_1() {
	fmt.Println(`
	var p1 Person will allocate an instance of Person
	p2 := new(Person) will allocate an Person instance and return an pointer to it
	`)
}

func Q20_2() {

	fmt.Println(`
	The 1st one will running wrong, it should be x := t and it will allocate a *T pointer and t will assign to it.
	The 2nd one will running wrong too, it should be x := &t and it will take address of t which in stacks and take it's pointer
	`)
}
