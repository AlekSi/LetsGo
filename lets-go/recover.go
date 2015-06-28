package main

import "fmt"

func main() {
	/*
		defer func() {
			if e := recover(); e != nil {
				fmt.Printf("recover: %v\n", e)
			}
		}()
	*/

	fmt.Println("Спокойно… Спокойно…")
	panic("Шеф, всё пропало!")
}
