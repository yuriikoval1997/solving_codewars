package main

import "fmt"

func main() {
	var arr1 = []int{4, 2, 8, 9, 11, 3}
	fmt.Println(MergeSort(arr1))
}

func MergeSort(array []int) []int {
	return divide(array)
}

func divide(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	var m = len(array) / 2
	var left = make([]int, m)
	var right = make([]int, len(array) - m)
	copy(left, array[:m])
	copy(right, array[m:])
	left = divide(left)
	right = divide(right)
	merge(left, right, array)
	return array
}

func merge(left []int, right []int, array []int) {
	var l = 0
	var r = 0
	var a = 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			array[a] = left[l]
			l++
		} else {
			array[a] = right[r]
			r++
		}
		a++
	}
	for l < len(left) {
		array[a] = left[l]
		l++
		a++
	}
	for r < len(right) {
		array[a] = right[r]
		r++
		a++
	}
}
