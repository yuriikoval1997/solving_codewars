package main

import "fmt"

func main() {
	var seconds int
	_, _ = fmt.Scan(&seconds)
	var hours = seconds / 3600
	var minutes = seconds / 60 % 60
	fmt.Printf("It is %d hours %d minutes.\n", hours, minutes)
}
