package main

import (
	"log"

	"github.com/bigwhite/gofmt-demo/pkg/bar"
	"github.com/bigwhite/gofmt-demo/pkg/foo"
)

func main() {
	foo.Foo()
	bar.Bar()
	log.Print("gofmt-demo")
}
