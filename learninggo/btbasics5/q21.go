package btbasics5

/*
Q21. (1) Linked List

1. Make use of the package container/list to create a (doubly) linked list.
Push the values 1, 2 and 4 to the list and then print it.

2. Create your own linked list implementation. And perform the same
actions as in question 1
*/

import "container/list"
import "fmt"

func Q21_1() {
	p := list.New()
	p.PushBack(1)
	p.PushBack(2)
	p.PushBack(3)
	p.PushBack(8)
	fmt.Printf("%v\n", p)

	for e := p.Front(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

//!!wang!! look at the implementation of go lang itself, it's good and fun.
func Q21_2() {

	p := New()
	p.Add(4)
	p.Add(5)
	p.Add(6)
	p.Add(7)

	for e := p.Head(); e != nil; e = e.Next() {
		fmt.Printf("%v ", e.Value)
	}
	fmt.Println()
}

type E struct {
	previous, next *E
	Value          int
}

type L struct {
	head *E
	end  *E
}

func New() *L {
	l := new(L)
	l.head = nil
	l.end = nil
	return l
}

func (l *L) Add(v int) int {
	if l.head == nil && l.end == nil {
		l.head = new(E)
		l.head.previous = nil
		l.head.next = nil
		l.head.Value = v
		l.end = l.head
		return v
	}
	temp := new(E)
	temp.Value = v
	l.end.next = temp
	temp.previous = l.end
	temp.next = nil
	l.end = temp

	return v
}

func (l *L) Head() *E {
	return l.head
}

func (e *E) Next() *E {
	return e.next
}
