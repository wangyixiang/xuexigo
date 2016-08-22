package main

import (
	"net/http"
	"io/ioutil"
	"time"
	"log"
	"fmt"
)

type Memo struct {
	f Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func NewMemo(f Func) *Memo {
	return &Memo{
		f:	f,
		cache:	make(map[string]result),
	}
}

func (memo *Memo) Get(key string) (interface{}, error)  {
	result, ok := memo.cache[key]
	if !ok {
		result.value, result.err = memo.f(key)
		memo.cache[key] = result
	}
	return result.value, result.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() []string{
	return []string {
		"www.baidu.com",
		"www.qq.com",
		"www.gopl.io",
		"www.baidu.com",
		"www.qq.com",
		"www.gopl.io",
	}
}

func main() {
	memo := NewMemo(httpGetBody)
	gstart := time.Now()
	for _, url := range incomingURLs() {
		start := time.Now()
		value, err := memo.Get("http://" + url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", "http://" + url, time.Since(start), len(value.([]byte)))
	}
	fmt.Println(time.Since(gstart))

}