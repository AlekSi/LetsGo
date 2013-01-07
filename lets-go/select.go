package main

import "fmt"

func main() {
	c1 := make(chan int, 5)
	c2 := make(chan int, 5)

	for i := 0; i < 5; i++ {
		c1 <- i
		c2 <- i
	}

	for {
		select {
		case v := <-c1:
			fmt.Printf("c1: %d\n", v)
		case v := <-c2:
			fmt.Printf("c2: %d\n", v)
		default:
			break
		}
	}
}
