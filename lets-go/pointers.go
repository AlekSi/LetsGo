package main

import "fmt"

type Conference struct {
	Name   string
	income int
}

func get() *Conference {
	conf := Conference{Name: "RailsClub'Ulyanovsk", income: 1}
	return &conf // HL
}

func main() {
	conf := get()
	fmt.Printf("%#v\n", conf)
	fmt.Printf("%#v\n", *conf)
}
