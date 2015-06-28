package main

import "fmt"

type Conference struct {
	Name   string
	Talks  int
	income float64 // доступно только в пакете main
}

func main() {
	conf := Conference{
		Name:   "РИТ++",
		Talks:  200,
		income: 100500,
	}
	conf.Name = "Golang-Moscow"
	fmt.Printf("%#v\n", conf)
}
