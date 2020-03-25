package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(T())
}

const (
	k = 1266
	p = 6
	v = 6
)

func T() float64 {
	return 6.0 / W()
}

func W() float64 {
	return math.Sqrt(k / M())
}

func M() float64 {
	return p * v
}
