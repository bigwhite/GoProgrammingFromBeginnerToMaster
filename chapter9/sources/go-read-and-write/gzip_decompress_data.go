package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "./hello_gopher.gz"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer f.Close()

	zw, _ := gzip.NewReader(f)
	defer zw.Close()

	i := 1
	for {
		buf := make([]byte, 32)
		_, err = zw.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("第%d次读取的压缩数据为: %q\n", i, buf)
				fmt.Println("读取到文件末尾，程序退出!")
			} else {
				fmt.Printf("第%d次读取压缩数据失败：%v", i, err)
			}
			return
		}
		fmt.Printf("第%d次读取的压缩数据为: %q\n", i, buf)
	}
}
