package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	pool := x509.NewCertPool()
	caCertPath := "../ca.crt"

	caCrt, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}
	pool.AppendCertsFromPEM(caCrt)

	s := &http.Server{
		Addr: "localhost:8081",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, World!\n")
		}),
		TLSConfig: &tls.Config{
			ClientCAs:  pool,
			ClientAuth: tls.RequireAndVerifyClientCert,
		},
	}

	fmt.Println(s.ListenAndServeTLS("../server-signed-by-ca.crt", "../server.key"))
}
