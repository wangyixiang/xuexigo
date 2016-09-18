package functions3

/*
Q9. (1) Stack
1. Create a simple stack which can hold a fixed number of ints. It does
not have to grow beyond this limit. Define push – put something on the
stack – and pop – retrieve something from the stack – functions. The
stack should be a LIFO (last in, first out) stack.
Figure 3.1. A simple LIFO stack
push(k)
pop() k
k
i
l
m
i++
i--
0

2. Bonus. Write a String method which converts the stack to a string rep-
resentation. This way you can print the stack using: fmt.Printf("My
stack %v\n", stack)
The stack in the figure could be represented as: [0:m] [1:l] [2:k]
*/
import "fmt"

type IntStack struct {
	stackp int
	stack  [50]int
}

func pop(s *IntStack) (errno, reti int) {
	if s.stackp == -1 {
		return -1, -1
	}
	reti = s.stack[s.stackp]
	s.stackp -= 1
	return 0, reti
}

func push(s *IntStack, value int) (errno int) {
	if s.stackp == (len(s.stack) - 1) {
		return -1
	}

	s.stackp += 1
	s.stack[s.stackp] = value
	return 0

}

func (s *IntStack) String() string {
	rs := ""
	for _, v := range s.stack {
		fmt.Printf("%d ", v)
		rs = rs + " " + string(v)
	}
	fmt.Println()
	return rs
}

func Q9_1() {
	ais := new(IntStack)
	push(ais, 1)
	push(ais, 2)
	push(ais, 'A')
	_, ai := pop(ais)
	fmt.Printf("Pop IntStack %d\n", ai)
	_, ai = pop(ais)
	fmt.Printf("Pop IntStack %d\n", ai)
	fmt.Print(ais)
}
