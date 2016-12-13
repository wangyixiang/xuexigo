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
	"strconv"
)

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

func outline1(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = stack[:len(stack) + 1]
		stack[len(stack) - 1] = n.Data // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline1(stack, c)
	}
}

var depth = 0

func pre1(node *html.Node) {
	if node.Type == html.ElementNode {
		depth ++
		fmt.Printf("%*s<%s>\n", depth * 2, " ", node.Data)
	}

}

func post1(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth * 2, " ", node.Data)
		depth --
	}
}

func outline21(node *html.Node, pre, post func(*html.Node)) {
	pre(node)

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		outline21(c, pre, post)
	}
	post(node)
}

func pre2(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth * 2, " ", node.Data)
	}

}

func post2(node *html.Node) {
	if node.Type == html.ElementNode {
		fmt.Printf("%*s</%s>\n", depth * 2, " ", node.Data)
	}
}

func outline22(node *html.Node, pre, post func(*html.Node)) {
	pre(node)

	if node.FirstChild != nil {
		depth ++
		outline22(node.FirstChild, pre, post)
		depth --
	}
	post(node)
	if node.NextSibling != nil {
		outline22(node.NextSibling, pre, post)
	}
}

func outline3(doc *html.Node) {
	f := func(node *html.Node) {
		stack := make([]string, 0)
		for node != nil {
			if node.Type == html.ElementNode {
				stack = append(stack, node.Data)
			} else {
				stack = append(stack, strconv.Itoa(int(node.Type)))
			}
			node = node.Parent
		}
		for i := len(stack) / 2 - 1; i >= 0; i-- {
			stack[i], stack[len(stack) - i - 1] = stack[len(stack) - i - 1], stack[i]
		}
		fmt.Println(stack)
	}
	visit(doc, f)

}

// 使用深度优先, 遍历整个html node tree
func visit(node *html.Node, f func(*html.Node)) {
	f(node)

	//for c := node.FirstChild; c != nil ; c = c.NextSibling {
	//	visit(c, f)
	//}

	if node.FirstChild != nil {
		visit(node.FirstChild, f)
	}

	if node.NextSibling != nil {
		visit(node.NextSibling, f)
	}
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
		filename = time.Now().Format("2006-01-02-15-04-05") + ".html"
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
		fmt.Println("outline under url:", url)
		outline(make([]string, 0, 128), node)
		outline1(make([]string, 0, 128), node)
		outline21(node, pre1, post1)
		outline22(node, pre2, post2)
		outline3(node)
		if *dumpFlag {
			err = dumpPage("", bs)
			if err != nil {
				log.Println(err)
			}
		}
		fmt.Println("-------------------------------------------------------------------------")
	}
}
