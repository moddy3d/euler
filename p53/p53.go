// Combinatoric selections
// https://projecteuler.net/problem=53

package main

import (
	"fmt"
	"math/big"
)

var (
	FactorialMap = map[int]*big.Int{}
)

func Factorial(n int) *big.Int {
	switch {
	case n > 1:
		f := big.NewInt(int64(n))
		f = f.Mul(f, Factorial(n-1))
		return f
	case n == 1:
		return big.NewInt(int64(1))
	default:
		return big.NewInt(int64(1))
	}
}

func Combinatoric(n, r int) *big.Int {
	f := Factorial(n)
	f = f.Div(f, Factorial(r))
	f = f.Div(f, Factorial(n-r))
	return f
}

func main() {
	over := 0
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			comb := Combinatoric(n, r)
			if comb.Cmp(big.NewInt(int64(1000000))) == 1 {
				over++
			}
		}
	}
	fmt.Println(over)
}
