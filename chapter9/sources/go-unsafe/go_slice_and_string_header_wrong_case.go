package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
	"unsafe"
)

func finalizer(p *[11]byte) {
	fmt.Println("数组对象被垃圾回收")
}

func newArray() *[11]byte {
	var a = [...]byte{'I', ' ', 'l', 'o', 'v', 'e', ' ', 'G', 'o', '!', '!'}
	runtime.SetFinalizer(&a, finalizer)
	return &a
}

func main() {
	var bh reflect.SliceHeader
	bh.Data = uintptr(unsafe.Pointer(newArray()))
	bh.Len = 11
	bh.Cap = 11
	var p = (*[]byte)(unsafe.Pointer(&bh))
	for i := 0; i < 3; i++ {
		runtime.GC()
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("%q\n", *p) //???
}
