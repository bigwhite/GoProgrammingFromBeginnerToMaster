package main

import "fmt"

func init() {
	fmt.Println("init invoked")
}

func main() {
	init()
}
