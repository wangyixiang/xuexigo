package main

import (
	"net/http"
	"log"
	"fmt"
	"crypto/x509"
	"io/ioutil"
	"os"
	"crypto/tls"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "你来啦!我的朋友!你可是经过了用户证书验证的人啊!我们共同的CA可以证明这一点.")
	})

	x509Pool := x509.NewCertPool()
	caPEM, err := ioutil.ReadFile("../ca/ca.crt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if !x509Pool.AppendCertsFromPEM(caPEM) {
		fmt.Fprintln(os.Stderr, "no any available Certificate Authority.")
		return
	}

	server := &http.Server{
		TLSConfig: &tls.Config{
			// 用来验证客户提供的证书是否是有合法CA签名的, 而目前这个Pool里面只有我提供的CA
			ClientCAs: x509Pool,
			// 这里不但要求客户通过Certificate, 而且我们还要对这个证书进行验证
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
		Addr:"0.0.0.0:443",
	}

	log.Fatal(server.ListenAndServeTLS("../ca/server.crt", "../ca/server.key"))
}
