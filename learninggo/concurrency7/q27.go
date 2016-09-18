package concurrency7

/*
Q27. (1) Channels
1. Modify the program you created in exercise Q2 to use channels, in other
words, the function called in the body should now be a goroutine and
communication should happen via channels. You should not worry
yourself on how the goroutine terminates.
2. There are a few annoying issues left if you resolve question 1. One
of the problems is that the goroutine isn’t neatly cleaned up when
main.main() exits. And worse, due to a race condition between the
exit of main.main() and main.shower() not all numbers are printed. It
should print up until 9, but sometimes it prints only to 8. Adding a sec-
ond quit-channel you can remedy both issues. Do this. a
*/

import "fmt"

//import "time"

var ch1 = make(chan int)

func Q27_1() {
	go q27_1()
	//time.Sleep(1000 * time.Millisecond)
	<-ch1
}

func q27_1() {
	for i := 1; i <= 10; i += 1 {
		fmt.Print(i)
	}
	fmt.Println()
	ch1 <- 1
}
