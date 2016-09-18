package interfaces6

/*
Q25. (1) Pointers and reflection
1. One of the last paragraphs in section “Introspection and reflection” on
page 205, has the following words:

The code on the right works OK and sets the member Name
to “Albert Einstein”. Of course this only works when you call
Set() with a pointer argument.
Why is this the case?

*/

import "fmt"

func Q25_1() {
	fmt.Println(`
	
	Listing 6.8. Reflect with private member
	
	type Person struct {
		name string "namestr" ← name
		age int
	}
	
	func Set(i interface{}) {
		switch i.(type) {
		case *Person:
			r := reflect.ValueOf(i)
			r.Elem(0).Field(0).SetString("
			Albert Einstein")
		}
	}
	
	Listing 6.9. Reflect with public member
	type Person struct {
		Name string "namestr" ← Name
		age int
	}
	
	func Set(i interface{}) {
		switch i.(type) {
		case *Person:
			r := reflect.ValueOf(i)
			r.Elem().Field(0).SetString("
			Albert Einstein")
		}
	}

	Because in Listing 6.8,
	name is not exported, so r.Elem(0) will not be allowed.
	
	Because in Listing 6.9
	"case *Person" which only have Person pointer case, so the Person
	will be not allowed.
	`)
}
