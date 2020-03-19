package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(Oper(Rot90Clock, "abcd\nefgh\nijkl\nmnop"))
	println()
	fmt.Println(Oper(Diag1Sym, "abcd\nefgh\nijkl\nmnop"))
	println()
	fmt.Println(Oper(SelfieAndDiag1, "abcd\nefgh\nijkl\nmnop"))
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

// Symmetry with respect to the main diagonal
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
