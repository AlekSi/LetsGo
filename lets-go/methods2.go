package main

import "fmt"

type Event struct {
	name string
}

func (c Event) Name() string { return fmt.Sprintf("Event %s", c.name) }

type Conference struct {
	Event
}

// func (c Conference) Name() string { return fmt.Sprintf("Conference %s", c.name) } // HL

func main() {
	conf := Conference{Event{"РИТ++"}}
	fmt.Printf("%s\n", conf.Name())
}
