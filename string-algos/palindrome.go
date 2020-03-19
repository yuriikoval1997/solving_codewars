package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	if strings.EqualFold(str, reverse(str)) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

func reverse(str string) string {
	var runes = []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
