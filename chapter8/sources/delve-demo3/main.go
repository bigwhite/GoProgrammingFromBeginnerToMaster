package main

import "fmt"

func main() {
	var p *int
	p = nil
	*p = 1
	fmt.Println("program exit")
}
