package main

import (
	"fmt"
	"sort"
)

type Event struct {
	Name string
}

type Events []Event

// START OMIT
func (c Events) Len() int {
	return len(c)
}

func (c Events) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}

func (c Events) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
	return
}

func main() {
	confs := Events{
		{"DevConf"},
		{"РИТ++"},
		{"Golang-Moscow"},
	}

	sort.Sort(confs)
	fmt.Println(confs)
}

// END OMIT
