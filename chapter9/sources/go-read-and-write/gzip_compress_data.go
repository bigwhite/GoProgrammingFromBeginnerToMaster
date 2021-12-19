package main

import (
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	file := "./hello_gopher.gz"
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer f.Close()

	zw := gzip.NewWriter(f)
	defer zw.Close()
	_, err = zw.Write([]byte("hello, gopher! I love golang!!"))
	if err != nil {
		fmt.Println("write compressed data error:", err)
	}
	fmt.Println("write compressed data ok")
}
