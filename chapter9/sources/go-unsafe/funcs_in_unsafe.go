package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	a int
	b string
	c [10]byte
	d float64
}

func main() {
	// fmt.Println(unsafe.Sizeof(int))
	var i int = 5
	var a = [100]int{}
	var sl = a[:]
	var f Foo

	fmt.Println(unsafe.Sizeof(i))           // 8
	fmt.Println(unsafe.Sizeof(a))           // 800
	fmt.Println(unsafe.Sizeof(sl))          // 24
	fmt.Println(unsafe.Sizeof(f))           // 48
	fmt.Println(unsafe.Sizeof(f.c))         // 10
	fmt.Println(unsafe.Sizeof((*int)(nil))) // 8

	fmt.Println(unsafe.Alignof(i))          // 8
	fmt.Println(unsafe.Alignof(f.a))        // 8
	fmt.Println(unsafe.Alignof(a))          // 8
	fmt.Println(unsafe.Alignof(sl))         // 8
	fmt.Println(unsafe.Alignof(f))          // 8
	fmt.Println(unsafe.Alignof(f.c))        // 1
	fmt.Println(unsafe.Alignof(struct{}{})) // 1
	fmt.Println(unsafe.Alignof([0]int{}))   // 8

	fmt.Println(unsafe.Offsetof(f.b)) // 8
	fmt.Println(unsafe.Offsetof(f.d)) // 40
}
