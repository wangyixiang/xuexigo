package main

import (
	"flag"
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
	"fmt"
	"time"
)

func walkDir(dir string, fileSize chan <- int64) {
	dirs := dEntries(dir)
	for _, entry := range dirs {
		if entry.IsDir() {
			walkDir(filepath.Join(dir, entry.Name()), fileSize)
			continue
		}
		fileSize <- entry.Size()
	}
}

func dEntries(dir string) []os.FileInfo {
	results, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("du1: %v\n", err)
		return nil
	}
	return results
}

func main() {
	var verbose = flag.Bool("v", false, "show spin when waiting.")
	flag.Parse()
	roots := flag.Args()
	var tick  <-chan time.Time
	if *verbose {
		tick = time.Tick(time.Millisecond * 500)
	}

	if len(roots) == 0 {
		roots = append(roots, ".")
	}
	chFileSize := make(chan int64)
	fmt.Println(time.Now())
	go func() {
		for _, root := range roots {
			walkDir(root, chFileSize)
		}
		close(chFileSize)
	}()
	nCount := 0
	nSize := int64(0)

	str := `-\|/`
	i := 0
	loop:
	for {
		select {
		case fSize, ok := <-chFileSize:
			if !ok {
				break loop
			}
			nCount += 1
			nSize += fSize
		case <-tick:
			fmt.Printf("\r%c%d files %dbytes %fG", str[i], nCount, nSize, float64(nSize) / 1e9)
			i = (i + 1) % 4
		}

	}
	fmt.Printf("\r%d files %dbytes %fG\n", nCount, nSize, float64(nSize) / 1e9)
	fmt.Println(time.Now())
}
