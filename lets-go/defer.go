package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("lets-go/defer.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// работа с f
}
