package main

import (
	"fmt"
)

func main() {
	const c = 99
	fmt.Println(c)

	fmt.Println(c + 1)

	const d int = 2
	fmt.Println(float32(d) + 1.9)
}
