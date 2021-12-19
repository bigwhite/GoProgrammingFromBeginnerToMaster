package main

import (
	"testing"
	"unsafe"
)

type A struct {
	a int8
	b [100][]byte
	c float32
}

type B struct {
	d int8
	e [100]string
	f float32
}

func newBFromA(a *A) *B {
	b := &B{
		d: a.a,
		f: a.c,
	}

	for i := 0; i < len(b.e); i++ {
		b.e[i] = string(a.b[i])
	}

	return b
}

func newBFromAByUnsafePointer(a *A) *B {
	b := &B{
		d: a.a,
		f: a.c,
	}
	for i := 0; i < len(b.e); i++ {
		b.e[i] = *(*string)((unsafe.Pointer)(&a.b[i]))
	}
	return b
}

func BenchmarkNewBFromA(b *testing.B) {
	var a = A{
		a: 17,
		b: [100][]byte{[]byte("hello, gophers! I love go!")},
		c: 3.1415,
	}
	for i := 0; i < b.N; i++ {
		b := newBFromA(&a)
		_ = b
	}
}

func BenchmarkNewBFromAByUnsafePointer(b *testing.B) {
	var a = A{
		a: 17,
		b: [100][]byte{[]byte("hello, gophers! I love go!")},
		c: 3.1415,
	}
	for i := 0; i < b.N; i++ {
		b := newBFromAByUnsafePointer(&a)
		_ = b
	}
}
