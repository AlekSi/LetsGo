package main

import "fmt"

type Conference struct {
	Name   string
	Talks  int
	income int
}

func get(name string) (conf *Conference, err error) {
	conf = &Conference{Name: name, Talks: 15, income: 10}
	return
}

func main() {
	conf, err := get("RailsClub'Ulyanovsk")
	fmt.Printf("%#v\n", conf)
	fmt.Printf("%#v\n", err)
}
