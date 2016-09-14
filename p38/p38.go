// https://projecteuler.net/problem=38

package main

import (
	"fmt"
	"strconv"
)

func Pandigital(n int) (int, bool) {
	digits := []rune{}
	runeMap := map[rune]bool{}
	for i := 1; i <= 9; i++ {
		product := i * n
		runes := []rune(strconv.Itoa(product))
		for _, r := range runes {
			if r == '0' {
				return 0, false
			}
			_, exists := runeMap[r]
			if exists {
				return 0, false
			}
			digits = append(digits, r)
			runeMap[r] = true
		}
		if len(digits) == 9 {
			pandigital, _ := strconv.Atoi(string(digits))
			return pandigital, true
		} else if len(digits) > 9 {
			return 0, false
		}
	}
	return 0, false
}

func main() {
	largest := 0
	for i := 192; i <= 10000; i++ {
		pan, isPan := Pandigital(i)
		if isPan && pan > largest {
			largest = pan
		}
	}
	fmt.Println(largest)
}
