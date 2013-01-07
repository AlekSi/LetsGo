package main

import "fmt"

type Conference struct {
	Name   string
	income int
}

func main() {
	confs := []Conference{
		Conference{"Стачка", 0},
		Conference{"RailsClub", 100500},
		Conference{"RailsClub'Ulyanovsk", 200},
	}
	for _, conf := range confs {
		fmt.Printf("%#v\n", conf)
	}
}
