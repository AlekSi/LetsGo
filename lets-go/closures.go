package main

import "fmt"

type Conference struct {
	Name   string
	Talks  int
	income float64
}

func main() {
	name := "РИТ++"
	get := func() (*Conference, error) { // HL
		return &Conference{Name: name, Talks: 200, income: 100500}, nil
	}
	conf, err := get()
	fmt.Printf("%#v\n", conf)
	fmt.Printf("%#v\n", err)
}
