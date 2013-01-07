package main

import "fmt"

type Event struct {
	Name string
}

type Conference struct {
	Event  // HL
	income int
}

func main() {
	conf := Conference{Event: Event{Name: "RailsClub"}}
	conf.Name = "RailsClub'Ulyanovsk"
	fmt.Printf("%#v\n", conf)
	fmt.Printf("%s\n", conf.Name) // HL
}
