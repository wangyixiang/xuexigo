package main

import (
	"time"
	"net/http"
	"fmt"
	"os"
)

func waitForServer(url string) error {
	const timeout = 30 * time.Second
	var err error
	sleepDuration := time.Second
	timeLimit := time.Now().Add(timeout)
	for time.Now().Before(timeLimit) {
		_, err = http.Get(url)
		if err != nil {
			time.Sleep(sleepDuration)
		} else {
			return err
		}
		sleepDuration = sleepDuration << 1
	}
	return fmt.Errorf("failed on visiting url:%s after %s with error: %v", url, timeout, err)
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(2)
	}
	if err := waitForServer(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stdout, "it's online now!\n")
	}
}
