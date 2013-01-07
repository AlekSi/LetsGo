package main

import "fmt"

func main() {
	array := [...]string{"0", "1", "2", "3", "4"}
	array[4] = "*"
	fmt.Printf("%#v\n\n", array)

	slice := array[1:]
	fmt.Printf("%#v\n\n", slice)

	slice = array[:]
	fmt.Printf("%#v\n\n", slice)
}
