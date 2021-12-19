package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file := "./bufio.txt"

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer f.Close()

	// 通过包裹函数创建带缓冲I/O的类型
	// 初始缓冲区大小为64字节
	bio := bufio.NewReaderSize(f, 64)
	fmt.Printf("初始状态下缓冲区缓存数据数量=%d字节\n\n", bio.Buffered())

	var i int = 1
	for {
		data := make([]byte, 15)
		n, err := bio.Read(data)
		if err == io.EOF {
			fmt.Printf("第%d次读取数据，读到文件末尾，程序退出\n", i)
			return
		}

		if err != nil {
			fmt.Println("读取数据出错：", err)
			return
		}

		fmt.Printf("第%d次读出数据：%q，长度=%d\n", i, data, n)
		fmt.Printf("当前缓冲区缓存数据数量=%d字节\n\n", bio.Buffered())
		i++
	}
}
