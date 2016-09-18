// Consecutive prime sum
// https://projecteuler.net/problem=50

package main

import (
	"fmt"
	"math"
)

func PrimesUnder(limit int) []bool {
	primes := make([]bool, limit+1, limit+1)
	for i := 0; i < limit; i++ {
		primes[i] = true
	}
	primes[0] = false
	primes[1] = false

	for i := 2; i < int(math.Sqrt(float64(limit))); i++ {
		if primes[i] {
			cache := i * i
			for k := 0; k < limit; k++ {
				j := cache + k*i
				if j > limit {
					break
				}
				primes[j] = false
			}
		}
	}

	return primes
}

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {
	primeMap := PrimesUnder(1000000)

	primes := []int{}
	for i, p := range primeMap {
		if p {
			primes = append(primes, i)
		}
	}

	most := 0
	for l := 0; l <= len(primes); l++ {
		sum := Sum(primes[l:len(primes)])
		for r := len(primes) - 1; r > l; r-- {
			sum -= primes[r]
			if sum < 1000000 {
				if primeMap[sum] {
					if r-l > most {
						most = r - l
						fmt.Println(sum, most)
					}
				}
			}
		}

	}
}
