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
		{"RailsClub'Ulyanovsk", 100},
	}
	for i := 0; i < len(confs); i++ {
		fmt.Printf("%#v\n", confs[i])
	}
}
