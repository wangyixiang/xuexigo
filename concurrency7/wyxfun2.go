package concurrency7

/*
from http://golang.org/doc/codewalk/sharemem/
I'm going to add an switch which to turn off the rolling service.

*/

import "fmt"
import "time"
import "net/http"
import "log"

//import "io/ioutil"

var results = make(map[string]string)

type status struct {
	url    string
	result string
}

func Wyxfun2_1() {
	var urls []string = []string{
		"http://www.google.com",
		"http://www.baidu.com",
		"http://weibo.com",
	}

	urlchan := make(chan *string)
	statuschan := make(chan *string)
	statusMonitor(urlchan, statuschan)
	go func() {
		var s *string
		for i := 0; i < 3; i++ {
			select {
			case s = <-statuschan:
				fmt.Print(s, "\n")
				fmt.Println(*s)
			}
		}
		close(urlchan)
	}()
	for _, aurl := range urls {
		urlchan <- &aurl
	}
	<-urlchan
}

func statusMonitor(url <-chan *string, status chan<- *string) {
	go func() {
		var checkurl *string
		var r *http.Response
		var err error
		for {
			select {
			case checkurl = <-url:
				fmt.Println(time.Now())
				r, err = http.Head(*checkurl)
				fmt.Println(time.Now())
				if err != nil {
					log.Println("Error ", *checkurl, err)
				}
				status <- &r.Status
			}
		}
	}()
}
