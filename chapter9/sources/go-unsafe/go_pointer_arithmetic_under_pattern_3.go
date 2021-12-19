package main

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	s string
	b int
	c float64
	d [10]int
}

func main() {
	var a = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var foo = Foo{
		s: "foo",
		b: 17,
		c: 3.1415,
		d: a,
	}

	var p = unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 3*unsafe.Sizeof(a[0]))
	fmt.Println(*(*int)(p))

	p = unsafe.Pointer(uintptr(unsafe.Pointer(&foo)) + unsafe.Offsetof(foo.c))
	fmt.Println(*(*float64)(p))

	for i := uintptr(0); i < unsafe.Sizeof((*int)(nil)); i++ {
		p = unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + i)
		fmt.Printf("0x%x\n", *(*byte)(p))
	}

	p = unsafe.Pointer(uintptr(unsafe.Pointer(&a)) + 10*unsafe.Sizeof(a[0])) // p已不在数组a的范围内了
	fmt.Println(*(*int)(p))
}
