#!/bin/bash
openssl genrsa -out ca.key 2048
#http://stackoverflow.com/questions/31506158/running-openssl-from-a-bash-script-on-windows-subject-does-not-start-with
openssl req -x509 -new -nodes -key ca.key -subj  "//CN=wangyixiang.com" -days 36500 -out ca.crt
openssl genrsa -out server.key 2048
openssl req -new -key server.key -subj "//CN=localhost" -out server.csr
openssl x509 -req -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days 3650
openssl genrsa -out client.key 2048
openssl req -new -key client.key -subj "//CN=wangyixiang.cn" -out client.csr
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days 3650
openssl x509 -req -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile client.ext -out client_ExtKeyUsage.crt -days 3650
openssl x509 -text -in client_ExtKeyUsage.crt -noout
