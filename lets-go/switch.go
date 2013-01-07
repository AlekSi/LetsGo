package main

import "fmt"

type Conference struct {
	Name   string
	income float32
}

func main() {
	conf := Conference{"RailsClub'Ulyanovsk", 400}
	switch conf.income {
	case 0:
		conf.income = 10000
		fallthrough // HL
	default:
		conf.income = conf.income * 0.9
	}
	fmt.Printf("%#v\n", conf)
}
