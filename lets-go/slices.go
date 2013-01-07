package main

import "fmt"

func main() {
	slice := []string{"0", "1", "2"}
	slice = append(slice, "3", "4")
	fmt.Printf("%#v\n\n", slice)

	slice = make([]string, 1, 2)
	fmt.Printf("%#v (len = %d, cap = %d)\n", slice, len(slice), cap(slice))

	slice = append(slice, "1")
	fmt.Printf("%#v (len = %d, cap = %d)\n", slice, len(slice), cap(slice))

	slice = append(slice, "2")
	fmt.Printf("%#v (len = %d, cap = %d)\n\n", slice, len(slice), cap(slice))

	slice = append(slice, "3", "4")
	slice = append(slice, slice...)
	fmt.Printf("%#v (len = %d, cap = %d)\n", slice, len(slice), cap(slice))
}
