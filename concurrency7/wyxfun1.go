package concurrency7

import (
	"fmt"
	"time"
)

var wyxch1 chan int
var wyxch2 chan int

func catcher(sleep int) {
	var count uint64 = 0
	for <-wyxch1 != 0 {
		count += 1
		time.Sleep(time.Duration(sleep) * time.Nanosecond)
	}

	fmt.Printf("I sleep %d nanoseconds, I cautch %d \n", sleep, count)
	wyxch2 <- sleep
}

func Wyxfun1_1() {
	catchers := []int{0, 100, 10000}
	wyxch1 = make(chan int)
	wyxch2 = make(chan int)
	for _, value := range catchers {
		go catcher(value)
	}
	for i := uint64(0); i <= 100000000; i++ {
		wyxch1 <- 1
	}
	close(wyxch1)
	<-wyxch2
	<-wyxch2
	<-wyxch2
}
