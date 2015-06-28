package main

import "fmt"

type Event struct {
	Name string
}

type Conference struct {
	Event  // HL
	income float64
}

func main() {
	conf := Conference{Event: Event{Name: "РИТ++"}}
	conf.Name = "Golang-Moscow"
	fmt.Printf("%#v\n", conf)
	fmt.Printf("%s\n", conf.Name) // conf.Event.Name // HL
}
