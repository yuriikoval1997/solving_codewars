package main

import "fmt"

func main() {
	var arr0 = []int{1, 2, 1}
	var arr1 = []int{0, 1, 2, 5, 1, 0}
	var arr2 = []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3}
	var arr3 = []int{1, 2, 2, 2, 1}
	var arr4 = []int{2, 1, 3, 1, 2, 2, 2, 2}
	var arr5 = []int{2, 1, 3, 1, 2, 2, 2, 2, 1}
	var arr6 = []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1}
	var arr7 = []int{-3, 12, 12, 9, -3, 14, 5, 1, -4}

	fmt.Println(PickPeaks(arr0))
	fmt.Println(PickPeaks(arr1))
	fmt.Println(PickPeaks(arr2))
	fmt.Println(PickPeaks(arr3))
	fmt.Println(PickPeaks(arr4))
	fmt.Println(PickPeaks(arr5))
	fmt.Println(PickPeaks(arr6))
	fmt.Println(PickPeaks(arr7))
}

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// If there is a plato, the algorithm shows the first element of it
// O(n) time complexity
func PickPeaks(array []int) PosPeaks {
	var pos = make([]int, 0)
	var peaks = make([]int, 0)
	if len(array) < 3 {
		return PosPeaks{pos, peaks}
	}
	var plato = false
	var platoStart int
	for i := 1; i < len(array)-1; i++ {
		if plato {
			if array[i] > array[i+1] {
				pos = append(pos, platoStart)
				peaks = append(peaks, array[platoStart])
				plato = false
			} else if array[i] < array[i+1] {
				plato = false
			}
		} else if array[i] > array[i-1] && array[i] > array[i+1] {
			pos = append(pos, i)
			peaks = append(peaks, array[i])
		} else if array[i] > array[i-1] && array[i] == array[i+1] {
			plato = true
			platoStart = i
		}
	}
	return PosPeaks{pos, peaks}
}
