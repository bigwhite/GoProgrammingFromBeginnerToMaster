package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a uint32 = 0x12345678
	fmt.Printf("0x%x\n", a)

	p := (unsafe.Pointer)(&a)

	b := (*[4]byte)(p)
	b[0] = 0x23
	b[1] = 0x45
	b[2] = 0x67
	b[3] = 0x8a

	fmt.Printf("0x%x\n", a)
}
