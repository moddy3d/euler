// Powerful digit counts
// https://projecteuler.net/problem=63

package main

import (
	"fmt"
	"math"
	"math/big"
)

func Power(n, p int64) *big.Int {
	power := big.NewInt(n)
	for p > 1 {
		power = power.Mul(power, big.NewInt(n))
		p--
	}
	return power
}

func main() {
	count := 0
	for n := 1; n < 10; n++ {
		p := 1
		for {
			power := Power(int64(n), int64(p))
			digits := len([]rune(power.String()))
			fmt.Println(n, p, power, digits)
			if digits == p {
				count++
			} else if math.Abs(float64(digits-p)) >= 2 {
				break
			}
			p++
		}
	}
	fmt.Println(count)
}
