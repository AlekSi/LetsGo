package main

import "fmt"

type Conference struct {
	Name   string
	income float64
}

func get() *Conference {
	conf := Conference{Name: "РИТ++", income: 100500}
	return &conf // так можно // HL
}

func main() {
	conf := get()
	fmt.Printf("%#v\n", conf)
	fmt.Printf("%#v\n", *conf)
}
