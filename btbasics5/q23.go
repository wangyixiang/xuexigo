package btbasics5

/*
Q23. (2) Method calls

1. Suppose we have the following program. Note the package contain-
er/vector was once part of Go, but has been removed when the append
built-in was introduced. However, for this question this isn’t impor-
tant. The package implemented a stack-like structure, with push and
pop methods.
package main
import "container/vector"func main() {
k1 := vector.IntVector{}
k2 := &vector.IntVector{}
k3 := new(vector.IntVector)
k1.Push(2)
k2.Push(3)
k3.Push(4)
}
What are the types of k1, k2 and k3?

2. Now, this program compiles and runs OK. All the Push operations work
even though the variables are of a different type. The documentation
for Push says:
func (p *IntVector) Push(x int) Push appends x to the end of
the vector.
So the receiver has to be of type *IntVector, why does the code above
(the Push statements) work correct then?

*/

import "fmt"
import "container/list"

func Q23_1() {
	fmt.Println("The package container/vector has been removed from the Go 1")
	fmt.Println("it's function have been replaced by built-in function, ")
	fmt.Println("append(), copy()")
	fmt.Println("check http://code.google.com/p/go-wiki/wiki/SliceTricks")

	fmt.Println("I use container/list to answer the question.")
	fmt.Println(`
	k1 := list.List{}
	k1 is a type List instance initialized with zero.
	k2 := &list.List{}
	k2 is a type List pointer instance.
	k3 := new(list.List)
	k3 is same as k2 , it's also a type List pointer instance.
	`)
}

func Q23_2() {
	fmt.Println(`
	http://golang.org/ref/spec#Calls
	The last paragraph of is section explained.
	A method call x.m() is valid if the method set of (the type of) x contains m and the argument list can be assigned to the parameter list of m. If x is addressable and &x's method set contains m, x.m() is shorthand for (&x).m(): 
	`)
	k1 := list.List{}
	k2 := &list.List{}
	k3 := new(list.List)

	k1.PushBack(1)
	k2.PushBack(2)
	k3.PushBack(3)
	list.New()
	q23_2()

}

func q23_2() int {
	return 1
}
