package main

import "fmt"

func main() {
	var arr1 = []int{0, 1, 2, 4, 5, 10, 11, 17}
	fmt.Println("Recursively")
	fmt.Println(FindRec(arr1, 1)) // 1
	fmt.Println(FindRec(arr1, -1)) // -1
	fmt.Println(FindRec(arr1, 4)) // 3
	fmt.Println(FindRec(arr1, 5)) // 4
	fmt.Println(FindRec(arr1, 17)) // 7
	fmt.Println(FindRec(arr1, 19)) // -1
	fmt.Println(FindRec(make([]int, 0), 19)) // -1

	fmt.Println("Iteratively")
	fmt.Println(FindIter(arr1, 1)) // 1
	fmt.Println(FindIter(arr1, -1)) // -1
	fmt.Println(FindIter(arr1, 4)) // 3
	fmt.Println(FindIter(arr1, 5)) // 4
	fmt.Println(FindIter(arr1, 17)) // 7
	fmt.Println(FindIter(arr1, 19)) // -1
	fmt.Println(FindIter(make([]int, 0), 19)) // -1
}

func FindRec(arr []int, n int) int {
	return search(arr, n, 0, len(arr)-1)
}

func search(arr []int, n int, s int, e int) int {
	if e-s < 0 {
		return -1
	}
	var m = (s + e) / 2
	if arr[m] == n {
		return m
	}
	if arr[m] > n {
		return search(arr, n, s, m-1)
	} else {
		return search(arr, n, m+1, e)
	}
}

func FindIter(arr []int, n int) int {
	var s = 0
	var e = len(arr) - 1
	for e >= s {
		var m = (s + e) / 2
		if arr[m] == n {
			return m
		}
		if arr[m] > n {
			e = m - 1
		} else {
			s = m + 1
		}
	}
	return -1
}
