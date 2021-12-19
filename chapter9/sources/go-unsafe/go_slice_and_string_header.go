package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func newSlice() *[]byte {
	var b = []byte("hello, gopher")
	return &b
}

func main() {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(newSlice()))
	var p = (*[]byte)(unsafe.Pointer(bh))
	fmt.Printf("%q\n", *p)

	var a = [...]byte{'I', ' ', 'l', 'o', 'v', 'e', ' ', 'G', 'o', '!', '!'}
	bh.Data = uintptr(unsafe.Pointer(&a))
	bh.Len = len(a)
	bh.Cap = len(a)

	fmt.Printf("%q\n", *p)
}
