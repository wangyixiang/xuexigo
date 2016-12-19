package main

import (
	"net/http"
	"fmt"
	"os"
	"io/ioutil"
	"crypto/x509"
	"crypto/tls"
)

func main() {
	x509pool := x509.NewCertPool()

	caPEM, err := ioutil.ReadFile("../ca/ca.crt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	if !x509pool.AppendCertsFromPEM(caPEM) {
		fmt.Fprintln(os.Stderr, "Failed on Append Certs From PEM")
		return
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			// 我们要对Server的Certificate进行验证, 所以我们在RootCAs加上俺制作的CA PEM
			RootCAs: x509pool,
		},
	}

	hsClient := &http.Client{
		Transport: transport,
	}

	resp, err := hsClient.Get("https://localhost")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	fmt.Fprintln(os.Stdout, string(data))
}

