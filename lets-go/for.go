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
		Conference{"RailsClub'Ulyanovsk", 100},
	}
	for i := 0; i < len(confs); i++ {
		fmt.Printf("%#v\n", confs[i])
	}
}
