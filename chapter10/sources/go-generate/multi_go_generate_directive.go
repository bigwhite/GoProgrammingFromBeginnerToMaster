//go:generate echo "top"
package main

import "fmt"

//go:generate echo "middle"
func main() {
	fmt.Println("hello, go generate")
}

//go:generate echo "tail"
