package main

import (
	"os"
	"io/ioutil"
	"log"
	"path/filepath"
	"flag"
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
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = append(roots, ".")
	}
	chFileSize := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, chFileSize)
		}
		close(chFileSize)
	}()
	nCount := 0
	nSize := int64(0)
	for fSize := range chFileSize {
		nCount += 1
		nSize += fSize
	}

	log.Printf("%d files %dbytes %fG", nCount, nSize,  float64(nSize)/1e9)
}