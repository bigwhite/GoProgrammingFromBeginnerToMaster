package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/bigwhite/delve-demo2/pkg/bar"
	"github.com/bigwhite/delve-demo2/pkg/foo"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			d := 2
			e := 20
			f := bar.Bar(d, e)
			fmt.Println(f)
			time.Sleep(2 * time.Second)
		}
		wg.Done()
	}()
	a := 3
	b := 10
	c := foo.Foo(a, b)
	fmt.Println(c)
	wg.Wait()
	fmt.Println("program exit")
}
