package main

import (
	"errors"
	"fmt"
)

func main() {
	var a int
	var b int
	_, _ = fmt.Scan(&a, &b)
	var res, err = divide(a, b)
	if err != nil {
		fmt.Println()
	} else {
		fmt.Println(res)
	}
}

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("ошибка")
	}
	return a / b, nil
}
