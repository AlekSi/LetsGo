package main

import "fmt"

type Event struct {
	name string
}

type Conference struct {
	Event
}

func (c Event) Name() string {
	return fmt.Sprintf("Event %s", c.name)
}

func (c Conference) Name() string {
	return fmt.Sprintf("Conference %s", c.name)
}

func main() {
	conf := Conference{Event{"RailsClub'Ulyanovsk"}}
	fmt.Printf("%s\n", conf.Name())
}
