package main

import "fmt"

func main() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Recover: %v\n", e)
		}
	}()
	panic("Шеф, всё пропало!")
}
