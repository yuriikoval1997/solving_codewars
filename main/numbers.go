package main

import (
	"errors"
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	fmt.Println("Simple way", DigitalRoot(n))
	for n > 9 {
		n = sumDigits(n, 0)
	}
	fmt.Println(n)
}

func sumDigits(n int, sum int) int {
	if n == 0 {
		return sum
	}
	return sumDigits(n / 10, sum + n%10)
}

func DigitalRoot(n int) int {
	return (n-1) % 9 + 1
}

func Max(numbers ...int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("cannot find maximum out of empty array")
	}
	var max = numbers[0]
	for i := 1; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return max, nil
}
