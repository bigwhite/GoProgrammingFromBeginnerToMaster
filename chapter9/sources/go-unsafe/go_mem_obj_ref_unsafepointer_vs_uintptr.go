package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"
)

type Foo struct {
	name string
}

func finalizer(p *Foo) {
	fmt.Printf("Foo: [%s]被垃圾回收\n", p.name)
}

func NewFoo(name string) *Foo {
	var f Foo = Foo{
		name: name,
	}
	runtime.SetFinalizer(&f, finalizer)
	return &f
}

func allocLargeObject() *[1000000]uint64 {
	a := [1000000]uint64{}
	return &a
}

func main() {
	var p1 = uintptr(unsafe.Pointer(NewFoo("FooRefByUintptr")))
	var p2 = unsafe.Pointer(NewFoo("FooRefByPointer"))

	for i := 0; i < 5; i++ {
		allocLargeObject()

		// 尝试输出p1和p2地址上的值
		q1 := (*Foo)(unsafe.Pointer(p1))
		fmt.Printf("object ref by uintptr: %+v\n", *q1)

		q2 := (*Foo)(p2)
		fmt.Printf("object ref by pointer: %+v\n", *q2)

		runtime.GC() // 运行垃圾回收
		time.Sleep(1 * time.Second)
	}
}
