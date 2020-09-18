package main

import "fmt"

func main() {
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3

	fmt.Println(arr1)
	fmt.Println(arr1[1])

	arr2 := [3]int{3, 4, 5}
	fmt.Println(arr2)
	fmt.Println(arr2[2])
}
