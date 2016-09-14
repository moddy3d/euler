// https://projecteuler.net/problem=36

package main

import (
	"fmt"
	"strconv"
)

func Palindrome(s string) bool {
	runes := []rune(s)
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-1-i] {
			return false
		}
	}
	return true
}

func DecimalToBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func main() {
	sum := 0
	for i := 0; i < 1000000; i++ {
		if Palindrome(DecimalToBinary(i)) && Palindrome(strconv.Itoa(i)) {
			sum += i
		}
	}
	fmt.Println(sum)
}
