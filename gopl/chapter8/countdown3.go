package main

import "fmt"

func main() {
	chCount := make(chan int)
	chAbort := make(chan int)
	for i := 1; i <= 10; i++ {
		select {
		case <-chCount:
			fmt.Println("chCount")
		case <-chAbort:
			fmt.Println("chAbort")
		default:
			fmt.Println(i)
		}
	}

	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x) // "0" "2" "4" "6" "8"
		case ch <- i:
		}
	}
}
