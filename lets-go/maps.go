package main

import "fmt"

func main() {
	m := map[string]int{"RailsClub": 10}
	fmt.Printf("%#v\n\n", m)

	m["RailsClub'Ulyanovsk"] = 15
	delete(m, "RailsClub")
	fmt.Printf("%#v\n\n", m)
}
