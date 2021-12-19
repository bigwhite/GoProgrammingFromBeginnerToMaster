package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var x = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("变量x的值=%d\n", x)
	println("变量x的地址=", q)

	var p = uintptr(unsafe.Pointer(&x))
	var q = unsafe.Pointer(&x)

	a(x) // 执行一系列函数调用

	// 变更数组x中元素的值
	for i := 0; i < 10; i++ {
		x[i] = x[i] + 10
	}

	println("栈扩容后，变量x的地址=", &x)
	fmt.Printf("栈扩容后，变量x的值=%d\n", x)

	fmt.Printf("变量p(uintptr)存储的地址上的值=%d\n", *(*[10]int)(unsafe.Pointer(p)))
	fmt.Printf("变量q(unsafe.Pointer)引用的地址上的值=%d\n", *(*[10]int)(q))
}

func a(x [10]int) {
	var y [100]int
	b(y)
}

func b(x [100]int) {
	var y [1000]int
	c(y)
}

func c(x [1000]int) {
}
