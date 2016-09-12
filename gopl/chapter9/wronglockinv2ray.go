package main

import (
	"log"
	"sync"
	"time"
)

// prove a wrong way of using Mutex

func lockInside(wg *sync.WaitGroup, n time.Duration) {
	var Finish sync.Mutex
	Finish.Lock()
	go func() {
		log.Println("Sleeping\t", int64(n), "Seconds")
		time.Sleep(n * time.Second)
		log.Println(time.Now())
		Finish.Unlock()
		wg.Done()
	}()
	Finish.Lock()
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go lockInside(wg, 2)
	wg.Add(1)
	go lockInside(wg, 1)
	wg.Wait()
}
