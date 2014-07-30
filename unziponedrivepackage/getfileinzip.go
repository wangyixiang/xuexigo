package unziponedrivepackage

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	indexFileName = "Encoding Errors.txt"
)

func getIndexMap(r io.Reader) (indexMap map[string]string, err error) {
	sStarter := "Original File Name  ->  New File Name"
	sMarker := "->"
	inputReader := bufio.NewReader(r)
	indexMap = make(map[string]string)

	for {
		line, err := inputReader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			return nil, err
		}
		if line == sStarter {
			break
		}
	}

	for {
		line, err := inputReader.ReadString('\n')
		if line == "" && err == nil {
			continue
		}
		iCount := strings.Count(line, sMarker)
		switch iCount {
		case 0:
			fmt.Println("Strange Data0:" + line)
		case 1:
			sSlice := strings.Split(line, sMarker)
			indexMap[strings.TrimSpace(sSlice[1])] = strings.TrimSpace(sSlice[0])
		default:
			fmt.Println("Strange Data2:" + line)
		}
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
	}
	return indexMap, err
}

func main() {
	f, ok := os.Open("testdata" + string(os.PathSeparator) + indexFileName)
	if ok == nil {
		im, _ := getIndexMap(f)
		fmt.Println(im)
	} else {
		log.Fatal(ok)
	}
}
