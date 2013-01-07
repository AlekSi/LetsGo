package main

import (
	"fmt"
	"sort"
)

type Conference struct {
	Name string
}

type Conferences []Conference

// START OMIT
func (c Conferences) Len() int {
	return len(c)
}

func (c Conferences) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}

func (c Conferences) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
	return
}

func main() {
	confs := Conferences{
		Conference{"Стачка"},
		Conference{"RailsClub"},
		Conference{"RailsClub'Ulyanovsk"},
	}

	sort.Sort(confs)
	fmt.Println(confs)
}

// END OMIT
