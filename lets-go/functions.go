package main

import "fmt"

type Conference struct {
	Name   string
	Talks  int
	income float64
}

func get(name string) (conf *Conference, err error) { // HL
	conf = &Conference{Name: name, Talks: 200, income: 100500}
	return // HL
}

func main() {
	conf, err := get("РИТ++")
	fmt.Printf("%#v\n", conf)
	fmt.Printf("%#v\n", err)
}
