package main

import (
	"fmt"
)

func main() {
	x := 1.0
	y := 2.0
	// mean = (x + y) / 2
	fmt.Printf("x=%v, of type %T\n", x, x)
	fmt.Printf("y=%v, of type %T\n", y, y)

	mean := (x + y) / 2
	fmt.Printf("result = %v, of type %T\n", mean, mean)
}
