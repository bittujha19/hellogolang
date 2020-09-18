package main

import (
	"fmt"
)

func main() {
	slice := []int{22, 33, 44}
	fmt.Println(slice)

	slice = append(slice, 55, 66, 77)
	fmt.Println(slice)

	s1 := slice[1:]
	s2 := slice[3:4]
	s3 := slice[:5]

	fmt.Println(s1, s2, s3)
}
