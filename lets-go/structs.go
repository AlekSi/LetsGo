package main

import "fmt"

type Conference struct {
	Name   string
	Talks  int
	income int // доступно только в пакете main
}

func main() {
	conf := Conference{Name: "RailsClub", Talks: 15, income: -1}
	conf.Name = "RailsClub'Ulyanovsk"
	fmt.Printf("%#v\n", conf)
}
