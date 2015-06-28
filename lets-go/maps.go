package main

import "fmt"

func main() {
	m := map[string]int{"РИТ++": 2000}
	fmt.Printf("%#v\n\n", m)

	m["Golang-Moscow"] = 514
	delete(m, "РИТ++")
	fmt.Printf("%#v\n\n", m)
}
