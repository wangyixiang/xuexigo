package interfaces6

/*
Q24. (1) Interfaces and compilation
1. The code in listing 6.3 on page 195 compiles OK — as stated in the text.
But when you run it you’ll get a runtime error, so something is wrong.
Why does the code compile cleanly then?
*/

import "fmt"

func Q24_1() {
	//even S doesn't explicitly declare that it implements I, but it's
	//a valid implementation of I
	fmt.Println(`
	type S struct {i int }
	func (p *S) Get() int { return p.i }
	func (p *S) Put(v int) { p.i = v }
	type I interface {
		Get() int
		Put(int)
	}
	
	func g(something interface{}) int {
		return something.(I).Get()
	}
	
	Listing 6.3
	i := 5
	fmt.Println(g(i))
	`)
	fmt.Println(`
	It will pass compilation, but int don't implement I,
	yes, it's a built-in int, it don't have method Get.
	!!wang!!
	My Question is compiler must already know that int don't have Get,
	why go lang choose panic at runtime but not compiling time?
	`)
}
