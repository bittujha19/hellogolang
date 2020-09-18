package main

import "fmt"

const (
	first = iota
	second
	third
	fourth
)

func main() {
	fmt.Println(first, second, third, fourth)
}
