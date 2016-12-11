package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
)

func fetchUrl(url string) ([]byte, error) {
	resp, err := http.Get(url)
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func displayPage(data []byte) {
	fmt.Fprintln(os.Stdout, string(data))
}

func dumpPage(data []byte) error {
	return ioutil.WriteFile("dumpedpage.txt", data, os.ModePerm)
}

func getUrlParameter() []string {
	flag.Parse()
	return flag.Args()
}

func main() {
	var dumpFlag = flag.Bool("d", false, "")
	urls := getUrlParameter()

	if len(urls) == 0 {
		fmt.Fprintln(os.Stderr, "give the site address after the command")
		return
	}
	for _, url := range urls {
		bs, err := fetchUrl(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error on fetching site:", url)
			continue
		}
		displayPage(bs)
		if *dumpFlag {
			err = dumpPage(bs)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error on dumpping site:", url)
			}
		}
	}
}