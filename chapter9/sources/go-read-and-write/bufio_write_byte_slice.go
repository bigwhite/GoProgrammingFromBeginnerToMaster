package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := "./bufio.txt"

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer func() {
		f.Sync()
		f.Close()
	}()

	data := []byte("I love golang!\n")

	// 通过包裹函数创建带缓冲I/O的类型
	bio := bufio.NewWriterSize(f, 32) //初始缓冲区大小为32字节

	// 将15个字节写入bio缓冲区，缓冲区缓存15个字节，bufio.txt中内容仍为空
	bio.Write(data)

	// 将15个字节写入bio缓冲区，缓冲区缓存30个字节, bufio.txt中内容仍为空
	bio.Write(data)

	// 将15个字节写入bio缓冲区后，bufio将32个字节写入bufio.txt，
	// bio缓冲区中仍然缓存15*3-32个字节
	bio.Write(data)

	// 将bio缓冲区中所有缓存数据均写入bufio.txt
	bio.Flush()
}
