package main

import "fmt"

type Conference struct {
	Name   string
	Talks  int
	income int
}

func main() {
	name := "RailsClub'Ulyanovsk"
	get := func() (conf *Conference, err error) {
		conf = &Conference{Name: name, Talks: 15, income: 500}
		return
	}
	conf, err := get()
	fmt.Printf("%#v\n", conf)
	fmt.Printf("%#v\n", err)
}
