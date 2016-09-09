package main

import (
	"fmt"
	"strings"
	"path"
)

func main() {
	apath := "王王/a/b/c/d/e//f.g.h"
	r := []rune(apath)
	fmt.Printf("%x\n%x\n", []byte(apath), r)
	fmt.Println(basename1(apath))
	fmt.Println(basename2(apath))
	fmt.Println(basename3(apath))
}

func basename1(path string) string {

	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '/' {
			path = path[i+1:]
			break
		}
	}

	for j := len(path) - 1; j >= 0; j-- {
		if path[j] == '.' {
			path = path[:j]
			break
		}
	}

	return path
}

func basename2(path string) string {
	path = path[strings.LastIndex(path, "/") + 1:]
	return path[:strings.LastIndex(path, ".")]
}

func basename3(path1 string) string {
	path1 = path.Base(path1)
	return path1[:strings.LastIndex(path1, ".")]
}

func basename4(path1 string) string {
	return ""
}