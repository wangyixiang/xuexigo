package main

import (
	"golang.org/x/net/html"
	"flag"
	"log"
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"os"
	"time"
)

func findlinks1(doc *html.Node) []string {
	links := make([]string, 0)

	links = visit(doc, links)

	return links
}

// 使用广度优先, 遍历整个html node tree
func visit(node *html.Node, links []string) []string {
	if node == nil {
		return links
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
				break
			}
		}
	}

	if node.NextSibling != nil {
		links = visit(node.NextSibling, links)
	}

	if node.FirstChild != nil {
		links = visit(node.FirstChild, links)

	}

	return links
}

func fetchUrl(url string) ([]byte, error) {
	request, _ := http.NewRequest("", url, nil)
	// 冒充Windows Chrome
	request.Header.Set("agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.71 Safari/537.36")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		log.Println("Opening ", url, " with StatusCode:", resp.StatusCode)
	}

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

func parseData(data []byte) (*html.Node, error) {
	buf := bytes.NewReader(data)
	node, err := html.Parse(buf)
	if err != nil {
		return nil, err
	}
	return node, nil
}

func dumpPage(filename string, data []byte) error {
	if filename == "" {
		filename = time.Now().Format("2006-01-02-15-04-05") + ".txt"
	}
	return ioutil.WriteFile(filename, data, os.ModePerm)
}

func main() {
	var dumpFlag = flag.Bool("d", false, "")
	flag.Parse()
	urls := flag.Args()
	for _, url := range urls {
		bs, err := fetchUrl(url)
		if err != nil {
			log.Println(err)
			continue
		}
		node, err := parseData(bs)
		if err != nil {
			log.Println(err)
			continue
		}
		links := findlinks1(node)
		fmt.Println("links under url:", url)
		if len(links) > 0 {
			for _, l := range links {
				fmt.Println(l)
			}
		}
		if *dumpFlag {
			err = dumpPage("", bs)
			if err != nil {
				log.Println(err)
			}
		}
		fmt.Println("-------------------------------------------------------------------------")
	}
}
