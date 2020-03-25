package main

import (
	"fmt"
	"strings"
)

var template = `a b c d
				e f g h
				i j k l
				m n o p`

func main() {
	fmt.Println("http://www.codewars.com/kata/56dbeec613c2f63be4000be6")
	fmt.Println(Oper(Rot90Clock, "abcd\nefgh\nijkl\nmnop"))
	println()
	fmt.Println(Oper(Diag1Sym, "abcd\nefgh\nijkl\nmnop"))
	println()
	fmt.Println(Oper(SelfieAndDiag1, "abcd\nefgh\nijkl\nmnop"))
	println()

	fmt.Println("https://www.codewars.com/kata/56dbf59b0a10feb08c000227/train/go")
	fmt.Println(Oper(Rot90Counter, "abcd\nefgh\nijkl\nmnop"))
	println()
	fmt.Println(Oper(Diag2Sym, "abcd\nefgh\nijkl\nmnop"))
	println()
	fmt.Println(Oper(SelfieDiag2Counterclock, "abcd\nefgh\nijkl\nmnop"))
}

// Clockwise rotation 90 degrees
func Rot90Clock(squareString string) string {
	var matrix = runeMatrix(squareString)
	var rowLength = len(matrix[0])
	var tmp = make([][]rune, len(matrix))
	for i := 0; i < rowLength; i++ {
		var str = make([]rune, rowLength)
		var k = 0
		for j := len(matrix) - 1; j >= 0; j-- {
			str[k] = matrix[j][i]
			k++
		}
		tmp[i] = str
	}
	var result = make([]string, len(tmp))
	for i, v := range tmp {
		result[i] = string(v)
	}
	return strings.Join(result, "\n")
}

// Symmetry with respect to the first main diagonal
func Diag1Sym(squareString string) string {
	var matrix = runeMatrix(squareString)
	var rowLength = len(matrix[0])
	var tmp = make([][]rune, len(matrix))
	for i := 0; i < len(matrix); i++ {
		var str = make([]rune, rowLength)
		var k = 0
		for j := 0; j < rowLength; j++ {
			str[k] = matrix[j][i]
			k++
		}
		tmp[i] = str
	}
	var result = make([]string, len(tmp))
	for i, v := range tmp {
		result[i] = string(v)
	}
	return strings.Join(result, "\n")
}

// It is initial string + string obtained by symmetry with respect to the main diagonal.
func SelfieAndDiag1(squareString string) string {
	var diag1 = strings.Split(Diag1Sym(squareString), "\n")
	var tmp = strings.Split(squareString, "\n")
	var result = make([]string, len(diag1))
	for i := 0; i < len(result); i++ {
		result[i] = strings.Join([]string{tmp[i], diag1[i]}, "|")
	}
	return strings.Join(result, "\n")
}

// Counterclockwise rotation 90 degrees
func Rot90Counter(squareString string) string {
	var matrix = runeMatrix(squareString)
	var rowLength = len(matrix[0])
	var tmp = make([][]rune, len(matrix))
	var k = 0
	for i := rowLength - 1; i >= 0; i-- {
		var str = make([]rune, rowLength)
		for j := 0; j < len(matrix); j++ {
			str[j] = matrix[j][i]
		}
		tmp[k] = str
		k++
	}
	var result = make([]string, len(tmp))
	for i := 0; i < len(result); i++ {
		result[i] = string(tmp[i])
	}
	return strings.Join(result, "\n")
}

// Symmetry with respect to the second main cross diagonal
func Diag2Sym(squareSequence string) string {
	var matrix = runeMatrix(squareSequence)
	var rowLength = len(matrix)
	var tmp = make([][]rune, len(matrix))
	var m = 0
	for i := rowLength - 1; i >= 0; i-- {
		var str = make([]rune, rowLength)
		var k = 0
		for j := rowLength - 1; j >= 0; j-- {
			str[k] = matrix[j][i]
			k++
		}
		tmp[m] = str
		m++
	}
	var result = make([]string, len(tmp))
	for i := 0; i < len(result); i++ {
		result[i] = string(tmp[i])
	}
	return strings.Join(result, "\n")
}

// It is initial string + string obtained by symmetry with respect
// to the main cross diagonal + counterclockwise rotation 90 degrees.
func SelfieDiag2Counterclock(squareString string) string {
	var self = strings.Split(squareString, "\n")
	var diag2 = strings.Split(Diag2Sym(squareString), "\n")
	var counterClock = strings.Split(Rot90Counter(squareString), "\n")
	var result = make([]string, len(self))
	for i := 0; i < len(self); i++ {
		result[i] = strings.Join([]string{self[i], diag2[i], counterClock[i]}, "|")
	}
	return strings.Join(result, "\n")
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

func runeMatrix(squareString string) [][]rune {
	var tmp = strings.Split(squareString, "\n")
	var rows = make([][]rune, len(tmp))
	for i, v := range tmp {
		rows[i] = []rune(v)
	}
	return rows
}
