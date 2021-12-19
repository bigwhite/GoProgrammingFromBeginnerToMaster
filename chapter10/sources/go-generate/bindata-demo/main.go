package main

import (
	"fmt"
	"net/http"
)

//go:generate go-bindata -o static.go static/img/go-mascot.jpg

func main() {
	data, err := Asset("static/img/go-mascot.jpg")
	if err != nil {
		fmt.Println("Asset invoke error:", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write(data)
	})

	http.ListenAndServe(":8080", nil)
}
