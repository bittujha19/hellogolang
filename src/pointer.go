package main

import (
	"fmt"
)

func main() {
	var i int
	i = 22
	fmt.Println(i)

	j := 44
	fmt.Println(j)

	var firstname = "bryan adams"
	fmt.Println(firstname)

	lastname := "Dyna hayden"
	fmt.Println(lastname)

	middlename := "James Anderson"
	ptr := &middlename
	fmt.Println(ptr, *ptr)

	var noname *string = new(string)
	*noname = "Nil Name"
	fmt.Println(noname, *noname)

	shopname := "mac donalds"
	ptr1 := &shopname
	fmt.Println(ptr1, *ptr1)

}
