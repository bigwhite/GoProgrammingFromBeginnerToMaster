package main

import (
	"fmt"

	"github.com/bigwhite/gorename-demo/pkg/foo"
)

func main() {
	var s = foo.PkgTypeStruct1{
		Field1: "tony",
		Age:    20,
	}
	s.MethodOfPkgType1()
	i := foo.NewPkgType1(5)
	fmt.Println(*i)
}
