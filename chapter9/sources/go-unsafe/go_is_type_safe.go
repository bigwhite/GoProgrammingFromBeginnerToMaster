package main

import "fmt"

func main() {
	a := 0x12345678
	fmt.Printf("0x%x\n", a)

	/*
		var p *byte = (*byte)(&a)
		*p = 0x23
	*/

	var b byte = byte(a)
	b = 0x23
	fmt.Printf("0x%x\n", b)
	fmt.Printf("0x%x\n", a)
}
