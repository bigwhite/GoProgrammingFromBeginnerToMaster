package main

import (
	"fmt"
	"io"
	"os"
)

func directWriteByteSliceToFile(path string, data []byte) (int, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return 0, err
	}
	defer func() {
		f.Sync()
		f.Close()
	}()

	return f.Write(data)
}

func directReadByteSliceFromFile(path string, data []byte) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error:", err)
		return 0, err
	}
	defer f.Close()

	return f.Read(data)
}

func main() {
	file := "./foo.txt"
	text := "hello, gopher"
	buf := make([]byte, 13)

	n, err := directWriteByteSliceToFile(file, []byte(text))
	if err != nil {
		fmt.Println("write file error:", err)
		return
	}
	fmt.Printf("write %d bytes to file.\n", n)

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer f.Close()
	for {
		n, err = f.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("read meets EOF")
				return
			}
			fmt.Println("read file error:", err)
		}
		fmt.Printf("read %d bytes from file: %q\n", n, buf)
	}
}
