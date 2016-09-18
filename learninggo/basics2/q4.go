// q4.go
package basics2

/*
1. Create a Go program that prints the following (up to 100 characters):64 Chapter 2: Basics
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...

2. Create a program that counts the number of characters in this string:
asSASA ddd dsjkdsjs dk
In addition, make it output the number of bytes in that string. Hint:
Check out the utf8 package.

3. Extend the program from the previous question to replace the three
runes at position 4 with ’abc’.

4. Write a Go program that reverses a string, so “foobar” is printed as
“raboof”. Hint: You will need to know about conversion; skip ahead to
section “Conversions” on page 164.

*/
import (
	"fmt"
	utf8 "unicode/utf8"
)

func Q4_1() {
	/*
		!!wang!!The way I answered this question, it may be bad, because I think it may
		generate the string type 99 times.
	*/
	a := "A"
	for i := 1; i <= 100; i++ {
		fmt.Println(a)
		a = a + "A"
	}
}

func Q4_2() {
	str := "asSASA ddd dsjkdsjs dk"
	runes := []rune(str)
	rcount := 0
	bcount := 0
	rcount = utf8.RuneCountInString(str)
	for i := 0; i < rcount; i++ {
		bcount += utf8.RuneLen(runes[i])
		fmt.Print(runes[i])
	}
	fmt.Println()
	fmt.Println(str)
	fmt.Println("rune in str is %d ", rcount)
	fmt.Println("len(str) is %d ", len(str))
	fmt.Println("byte in str is %d ", bcount)

}

func Q4_3() {
	str := "asSASA ddd dsjkdsjs dk"
	fmt.Println(str)
	runes := []rune(str)
	runes[3] = 'a'
	runes[4] = 'b'
	runes[5] = 'c'
	fmt.Println("Replace three runes at position 4 with \"abc\"")
	fmt.Println(string(runes))
}

func Q4_4() {
	fmt.Println(q4_4("go xuexi bu rongyi!!!"))
}

func q4_4(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
