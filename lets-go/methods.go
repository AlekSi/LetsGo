package main

import "fmt"

type Conference struct {
	Name   string
	income float32
}

func (c Conference) Income() float32 {
	return c.income * 0.9
}

func main() {
	conf := Conference{Name: "RailsClub'Ulyanovsk", income: 1000}
	fmt.Printf("%f\n", conf.Income())
}
