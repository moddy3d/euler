// Ordered fractions
// https://projecteuler.net/problem=71

package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		Max  int64    = 1000000
		a    int64    = 3
		b    int64    = 7
		left *big.Rat = nil
	)

	target := big.NewRat(a, b)
	minDiff := big.NewRat(int64(1), int64(1))
	for d := b + 1; d < Max; d++ {
		n := (float64(d) * float64(a)) / float64(b)
		r := big.NewRat(int64(n), int64(d))
		if r.Cmp(target) == -1 {
			s := big.NewRat(int64(1), int64(1)).Sub(target, r)
			if minDiff.Cmp(s) == 1 {
				minDiff = s
				left = r
			}
		}
	}
	fmt.Println(left)
}
