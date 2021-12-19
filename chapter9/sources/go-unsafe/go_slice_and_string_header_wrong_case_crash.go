//go run -gcflags="-l" go_slice_and_string_header_wrong_case_crash.go
package main

import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

func finalizer(p *[66]byte) {
	fmt.Println("数组对象被垃圾回收")
}

func allocLargeObject() *[1000000]uint64 {
	a := [1000000]uint64{}
	return &a
}

func newArray() *[66]byte {
	var a = [...]byte{'I', ' ', 'l', 'o', 'v', 'e', ' ', 'G', 'o', '!', '!',
		'I', ' ', 'l', 'o', 'v', 'e', ' ', 'G', 'o', '!', '!',
		'I', ' ', 'l', 'o', 'v', 'e', ' ', 'G', 'o', '!', '!',
		'I', ' ', 'l', 'o', 'v', 'e', ' ', 'G', 'o', '!', '!',
		'I', ' ', 'l', 'o', 'v', 'e', ' ', 'G', 'o', '!', '!',
		'I', ' ', 'l', 'o', 'v', 'e', ' ', 'G', 'o', '!', '!'}
	runtime.SetFinalizer(&a, finalizer)
	return &a
}

func main() {
	var bh reflect.SliceHeader
	bh.Data = uintptr(unsafe.Pointer(newArray()))
	bh.Len = 66
	bh.Cap = 66
	var p = (*[]byte)(unsafe.Pointer(&bh))
	for i := 0; i < 100; i++ {
		runtime.GC()
		allocLargeObject()
		fmt.Printf("%q\n", *p) //???
	}
}
