// https://projecteuler.net/problem=34

package main

import (
	"fmt"
	"strconv"
)

var (
	FactorialMap = map[int]int{}
)

func Factorial(n int) int {
	switch {
	case n > 1:
		f, exists := FactorialMap[n]
		if exists {
			return f
		}

		f = n * Factorial(n-1)
		FactorialMap[n] = f
		return f
	case n == 1:
		return 1
	default:
		return 1
	}
}

func FactorialSum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += Factorial(n)
	}
	return sum
}

func NumberToDigits(n int) []int {
	digits := []int{}
	runes := strconv.Itoa(n)
	for _, r := range runes {
		digit, _ := strconv.Atoi(string(r))
		digits = append(digits, digit)
	}
	return digits
}

func main() {
	sum := 0
	limit := Factorial(9) * 7
	fmt.Println(limit)
	for n := 10; n < limit; n++ {
		if n == FactorialSum(NumberToDigits(n)) {
			sum += n
		}
	}
	fmt.Println(sum)
}
