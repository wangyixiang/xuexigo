package btbasics5

/*
Q22. (1) Cat

1. Write a program which mimics the Unix program cat. For those who
don’t know this program, the following invocation displays the con-
tents of the file blah:
% cat blah

2. Make it support the n flag, where each line is numbered.
*/

import "fmt"
import "os"

func Q22_1(path string) {
	fp, _ := os.Open(path)
	buf := make([]byte, 1024)
	n, _ := fp.Read(buf)
	for n != 0 {
		fmt.Printf("%v ", string(buf[:n]))
		n, _ = fp.Read(buf)

	}
	fp.Close()
}

func Q22_2(path string) {
	fp, _ := os.Open(path)
	buf := make([]byte, 1024)
	n, _ := fp.Read(buf)
	linecount := 1
	for n != 0 {
		startm := 0
		endm := 0
		for i, value := range buf {
			if i == n {
				break
			}
			if value == '\n' {
				endm = i + 1
				fmt.Printf("%d: ", linecount)
				fmt.Printf("%v", string(buf[startm:endm]))
				linecount += 1
				startm = endm
			}
		}
		n, _ = fp.Read(buf)
	}
}
