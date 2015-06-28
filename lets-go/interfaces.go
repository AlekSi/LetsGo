package main

import "fmt"

type Event interface {
	Name() string
}

type RIT struct{}

func (c RIT) Name() string {
	return "РИТ++"
}

func main() {
	conf := new(RIT) // = &RIT{}
	fmt.Println(conf.Name())
}
