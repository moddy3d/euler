// Goldbach's other conjecture
// https://projecteuler.net/problem=46

package main

import (
	"fmt"
	"math"
)

func PrimesUnder(limit int) []bool {
	numbers := make([]bool, limit+1, limit+1)
	for i := 0; i < limit; i++ {
		numbers[i] = true
	}
	numbers[0] = false
	numbers[1] = false

	for i := 2; i < int(math.Sqrt(float64(limit))); i++ {
		if numbers[i] == true {
			cache := i * i
			for k := 0; k < limit; k++ {
				j := cache + k*i
				if j > limit {
					break
				}
				numbers[j] = false
			}
		}
	}

	return numbers
}

func Goldbach(num int, primes *[]bool) bool {
	for prime := num - 1; prime >= 2; prime-- {
		if (*primes)[prime] {
			twoSquare := num - prime
			if math.Mod(float64(twoSquare), 2.0) == 0 {
				square := twoSquare / 2
				sqrt := math.Sqrt(float64(square))
				if sqrt == float64(int64(sqrt)) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	primes := PrimesUnder(1000000)
	for i := 9; i < len(primes); i++ {
		if !primes[i] && !Goldbach(i, &primes) && math.Mod(float64(i), 2.0) != 0.0 {
			fmt.Println(i)
			break
		}
	}
}
