package main

import "fmt"

func main() {
	var nums = []float32{1, 1, 1, 2, 1, 1}
	fmt.Println(findUniqueNumber(nums))
}

func findUniqueNumber(arr []float32) float32 {
	if len(arr) < 3 {
		panic("cannot find unique value if there are less then 3 elements")
	}
	var x = arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] != x {
			if arr[i] != arr[i+1] {
				return arr[i]
			} else {
				return x
			}
		}
	}
	if arr[len(arr) - 1] == arr[len(arr) - 2] {
		panic("there is no unique elements")
	}
	return arr[len(arr)-1]
}
