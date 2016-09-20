// Lychrel numbers
// https://projecteuler.net/problem=55

package main

import (
	"fmt"
	"math/big"
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

func Reverse(s string) string {
	runes := []rune(s)
	reverseRunes := []rune{}
	for i := len(runes) - 1; i >= 0; i-- {
		reverseRunes = append(reverseRunes, runes[i])
	}
	return string(reverseRunes)
}

func Lychrel(num *big.Int, iter int) bool {
	if iter >= 50 {
		return true
	}

	s := num.String()
	r := Reverse(s)
	reverseNum := new(big.Int)
	reverseNum, success := reverseNum.SetString(r, 10)
	if !success {
		fmt.Printf("Could not set string %s\n", r)
	}

	num = num.Add(num, reverseNum)
	if Palindrome(num.String()) {
		return false
	} else {
		return Lychrel(num, iter+1)
	}
}

func main() {
	count := 0
	for i := 0; i < 10000; i++ {
		if Lychrel(big.NewInt(int64(i)), 1) {
			count++
		}
	}
	fmt.Println(count)
}
