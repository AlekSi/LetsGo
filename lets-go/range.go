package main

import "fmt"

type Conference struct {
	Name   string
	income float64
}

func main() {
	confs := []Conference{
		{"Стачка", 0},
		{"RailsClub", 100500},
		{"RailsClub'Ulyanovsk", 200},
	}
	for _, conf := range confs {
		fmt.Printf("%#v\n", conf)
	}
}
