package main

import "fmt"

type Conference struct {
	Name   string
	income float64
}

func main() {
	conf := Conference{"РИТ++", 100500}
	switch conf.income {
	case 0:
		conf.income = 100000
		fallthrough // HL
	default:
		conf.income = conf.income * 0.9
	}
	fmt.Printf("%#v\n", conf)
}
