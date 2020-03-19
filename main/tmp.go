package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	var str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	var pattern = regexp.MustCompile("\\s")
	str = pattern.ReplaceAllString(str, "s")
	if strings.EqualFold(str, reverse(str)) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}

func reverse(str string) string {
	var runes = []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes) && j >= 0; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}