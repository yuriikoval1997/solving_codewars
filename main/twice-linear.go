package main

import (
	"fmt"
)

func main() {
	fmt.Println(DblLinear(10))
}

// For each x in seq, then y = 2 * x + 1 and z = 3 * x + 1 must be in seq too.
//
// Given parameter n the function dbl_linear (or dblLinear...)
// returns the element seq(n) of the ordered (with <) sequence seq (so, there are no duplicates).
func DblLinear(n int) int {
	var seq = make([]int, 0, n+1)
	seq = append(seq, 1)
	var i = 0
	var j = 0
	for len(seq) <= n {
		var yi = 2 * seq[i] + 1
		var zj = 3 * seq[j] + 1
		if yi < zj {
			seq = append(seq, yi)
			i++
		} else if yi > zj {
			seq = append(seq, zj)
			j++
		} else {
			seq = append(seq, yi)
			i++
			j++
		}
	}
	return seq[n]
}
