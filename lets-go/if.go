package main

import "fmt"

type Conference struct {
	Name   string
	income float64
}

func main() {
	confs := []Conference{
		{"DevConf", 9000},
		{"РИТ++", 100500},
		{"Golang-Moscow", 0},
	}
	for _, conf := range confs { // HL
		if conf.income >= 10000 {
			conf.income = conf.income * 0.8
		} else {
			conf.income = conf.income * 0.9
		}
		fmt.Printf("%#v\n", conf)
	}
}
