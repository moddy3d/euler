// Powerful digit sum
// https://projecteuler.net/problem=56

package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func Length(i *big.Int) int {
	s := i.String()
	sum := 0
	for _, r := range []rune(s) {
		val, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		sum += val
	}
	return sum
}

func main() {
	greatest := 0
	for a := 2; a < 100; a++ {
		for b := 1; b < 100; b++ {
			i := big.NewInt(int64(a))
			i = i.Exp(i, big.NewInt(int64(b)), nil)
			sum := Length(i)
			if sum > greatest {
				greatest = sum
			}
		}
	}
	fmt.Println(greatest)
}
