package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = [10]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	var p = (*uint64)(unsafe.Pointer(&a[5]))

	fmt.Printf("a[5]的地址: 0x%p\n", &a[5])
	fmt.Printf("a[5]的对齐系数: %d\n", unsafe.Alignof(a[5]))
	fmt.Printf("p的地址: 0x%p\n", p)
	fmt.Printf("p的对齐系数: %d\n", unsafe.Alignof(p))
	fmt.Printf("*p = %d\n", *p)
}
