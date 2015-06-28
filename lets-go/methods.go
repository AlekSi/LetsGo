package main

import "fmt"

type Conference struct {
	Name   string
	income float64
}

func (c Conference) Income() float64 { // или (c *Conference) // HL
	return c.income * 0.9
}

func main() {
	conf := Conference{Name: "РИТ++", income: 100500}
	fmt.Printf("%f\n", conf.Income())
}
