// Square root convergents
// https://projecteuler.net/problem=57

package main

import (
	"fmt"
	"math/big"
)

var (
	ExpandCache = map[int]*big.Rat{}
)

func Length(n *big.Int) int {
	return len([]rune(n.String()))
}

func Expand(depth int) *big.Rat {
	result, exists := ExpandCache[depth]
	if exists {
		return result
	}

	a := big.NewRat(int64(2), int64(1))
	if depth == 0 {
		b := big.NewRat(int64(1), int64(2))
		a = a.Add(a, b)
		return a
	} else {
		b := big.NewRat(int64(1), int64(1))
		c := Expand(depth - 1)
		b = b.Quo(b, c)
		a = a.Add(a, b)
		ExpandCache[depth] = a
		return a
	}
}

func main() {
	count := 0
	for n := 0; n < 1000; n++ {
		e := Expand(n)
		a := big.NewRat(int64(1), int64(1))
		b := big.NewRat(int64(1), int64(1))
		a = a.Quo(a, e)
		b = b.Add(b, a)
		if Length(b.Num()) > Length(b.Denom()) {
			count++
		}
	}
	fmt.Println(count)
}
