package main

import "fmt"

type Event interface {
	Name() string
}

type Conference struct{}

func (c Conference) Name() string {
	return "RailsClub'Ulyanovsk"
}

func main() {
	conf := new(Conference) // = &Conference{}
	fmt.Println(conf.Name())
}
