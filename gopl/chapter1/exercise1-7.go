package main

import (
	"os"
	"log"
	"strings"
	"net/http"
	"io"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("give me a site address.")
	}
	siteUrl := os.Args[1]
	if !strings.HasPrefix(siteUrl, "https://") && !strings.HasPrefix(siteUrl, "http://") {
		siteUrl = "http://" + siteUrl
	}
	resq, err := http.Get(siteUrl)
	if err != nil {
		log.Fatal(err)
	}
	io.Copy(os.Stdout, resq.Body)
}
