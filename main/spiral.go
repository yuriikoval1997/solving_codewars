package main

import "fmt"

func main() {
	fmt.Println("5x5")
	var spiral = CreateSpiral(5)
	for i := 0; i < len(spiral); i++ {
		fmt.Println(spiral[i])
	}

	fmt.Println("6x6")
	spiral = CreateSpiral(6)
	for i := 0; i < len(spiral); i++ {
		fmt.Println(spiral[i])
	}

	fmt.Println("-1")
	spiral = CreateSpiral(-1)
	for i := 0; i < len(spiral); i++ {
		fmt.Println(spiral[i])
	}
}

func CreateSpiral(n int) [][]int {
	if n < 0 {
		return [][]int{}
	}
	var spiral = make([][]int, n)
	for i := 0; i < n; i++ {
		spiral[i] = make([]int, n)
	}
	var value = 1
	var startRow = 0
	var endRow = n
	var startColumn = 0
	var endColumn = n
	for startRow < endRow && startColumn < endColumn {
		for i := startColumn; i < endColumn; i++ {
			spiral[startRow][i] = value
			value++
		}
		startRow++
		for i := startRow; i < endRow; i++ {
			spiral[i][endColumn-1] = value
			value++
		}
		endColumn--
		for i := endColumn - 1; i >= startColumn; i-- {
			spiral[endRow-1][i] = value
			value++
		}
		endRow--
		for i := endRow - 1; i >= startRow; i-- {
			spiral[i][startColumn] = value
			value++
		}
		startColumn++
	}
	return spiral
}
