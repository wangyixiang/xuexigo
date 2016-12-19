package main

import (
	"net/http"
	"log"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "你来啦!")
	})
	log.Fatal(http.ListenAndServeTLS("0.0.0.0:443", "../ca/server.crt", "../ca/server.key", nil))
}
