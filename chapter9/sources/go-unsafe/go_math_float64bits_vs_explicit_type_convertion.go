package main

import (
	"fmt"
	"math"
)

func main() {
	var f float64 = 3.1415
	var d1 = uint64(f)
	var d2 = math.Float64bits(f)
	fmt.Printf("d1 = %d, d2 = %d\n", d1, d2)
}
