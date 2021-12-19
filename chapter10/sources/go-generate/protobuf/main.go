package main

import (
	"fmt"

	msg "github.com/bigwhite/protobuf-demo/msg"
)

//go:generate protoc -I ./IDL msg.proto --gofast_out=./msg
func main() {
	var m = msg.Request{
		MsgID:  "xxxx",
		Field1: "field1",
		Field2: []string{"field2-1", "field2-2"},
	}
	fmt.Println(m)
}
