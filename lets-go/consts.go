package main

import (
	"fmt"
)

func main() {
	const pi = 3.14159265358979323846264338327950288419716939937510582097494459
	var pi32 float32 = pi
	pi64 := float64(pi)
	fmt.Printf("%v\n%v\n", pi32, pi64)
}
