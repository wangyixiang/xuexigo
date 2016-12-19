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
		fmt.Fprintln(os.Stderr, "Failed on Append Certs From PEM.")
		return
	}

	clientCrt, err := tls.LoadX509KeyPair("../ca/client.crt", "../ca/client.key")

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed on Load client key pair.")
	}

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			RootCAs: x509pool,
			// 因为服务器端要要求并会验证我们的证书, 所以我们就只好乖乖地把证书奉上, 当然, private key是不会给的, 是用来
			// 本地加密计算的.
			Certificates: []tls.Certificate{clientCrt},
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

