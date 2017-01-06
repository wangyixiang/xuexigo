package main

import (
	"sync"
	"time"
	"fmt"
)

type rwLock struct {
	sync.RWMutex
	count int
}

func main() {
	wg := &sync.WaitGroup{}

	rwlock := &rwLock{}

	wg.Add(1)
	go func() {
		rwlock.RLock()
		time.Sleep(15*time.Second)
		fmt.Println("reading goroutine 1 sleep 15s, count is ", rwlock.count)
		wg.Done()
		rwlock.RUnlock()
	}()

	wg.Add(1)
	go func() {
		rwlock.RLock()
		time.Sleep(5*time.Second)
		fmt.Println("reading goroutine 2 sleep 5s, count is ", rwlock.count)
		wg.Done()
		rwlock.RUnlock()
	}()

	wg.Add(1)
	go func() {
		rwlock.Lock()
		time.Sleep(10*time.Second)
		fmt.Println("writing goroutine 1 sleep 10s")
		rwlock.count++
		wg.Done()
		rwlock.Unlock()
	}()

	wg.Add(1)
	go func() {
		rwlock.Lock()
		time.Sleep(5*time.Second)
		fmt.Println("writing goroutine 2 sleep 5s")
		rwlock.count++
		wg.Done()
		rwlock.Unlock()
	}()

	wg.Add(1)
	go func() {
		time.Sleep((15 + 10 + 3)*time.Second)
		rwlock.RLock()
		fmt.Println("reading goroutine 3 should happen after 30s and count should be 2 and it'is ", rwlock.count)
		wg.Done()
		rwlock.RUnlock()
	}()

	ticker := time.NewTicker(time.Second)
	count := 0

	go func() {
		for {
			select {
			case <-ticker.C:
				count += 1
				fmt.Printf("\r%ds ", count)
			}
		}
	}()

	wg.Wait()
	fmt.Printf("\r%d", count)
}
