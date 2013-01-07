package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("lets-go/defer.go")
	fmt.Printf("Error: %v\n", err)
	defer f.Close()

	// работа с f
}
