package main

import "fmt"

func main() {
	strings := []string{"a","","data"}
	fmt.Println(strings)
	noempty1(strings)
	fmt.Println(strings)
	strings = []string{"a","","data"}
	fmt.Println(strings)
	noempty2(strings)
	fmt.Println(strings)
}

func noempty1(strings []string) []string {
	i := 0
	for _, str := range strings {
		if str != "" {
			strings[i] = str
			i += 1
		}
	}
	return strings[:i]
}

func noempty2(strings []string) []string {
	result := strings[:0]
	for _, str := range strings {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}