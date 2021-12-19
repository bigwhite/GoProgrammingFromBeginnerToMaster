package main

import (
	"fmt"

	"github.com/bigwhite/delve-demo1/pkg/foo"
)

func main() {
	a := 3
	b := 10
	c := foo.Foo(a, b)
	fmt.Println(c)
}
