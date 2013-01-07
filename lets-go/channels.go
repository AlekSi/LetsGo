package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	for i := 0; i < 5; i++ {
		c <- i
	}
}
