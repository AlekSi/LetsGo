package main

import (
	"fmt"
	"time"
)

func count(g int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d: %d\n", g, i)
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	go func() {
		count(1)
	}()
	count(2)
}
