// Totient permutation
// https://projecteuler.net/problem=70

package main

import (
	"fmt"
	"math"
)

var (
	Max = 10000000
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

func LargestPermutation(n int) int {
	digits := make([]int, 10)
	perm := 0

	k := n
	for k > 0 {
		digits[int(math.Mod(float64(k), 10.0))]++
		k /= 10
	}

	for i := 9; i >= 0; i-- {
		for j := 0; j < digits[i]; j++ {
			perm = perm*10 + i
		}
	}

	return perm
}

func main() {
	// Generate primes
	primeMap := PrimesUnder(Max)
	primes := []int{}
	for n, isPrime := range primeMap {
		if isPrime {
			primes = append(primes, n)
		}
	}

	// Generate totients
	totients := make([]int, Max+1, Max+1)
	for n := 0; n < Max+1; n++ {
		totients[n] = n
	}
	totients[0] = 1
	totients[1] = 1

	for _, n := range primes {
		multiple := 1
		for {
			factor := n * multiple
			if factor < Max+1 {
				if factor == 10 {
					fmt.Println(n, multiple, factor)
				}
				if primeMap[multiple] && multiple > n {
					totients[factor] -= multiple - 1
				} else {
					totients[factor] -= multiple
				}
			} else {
				break
			}
			multiple++
		}
	}

	// Find Ratio
	fmt.Println(totients[2:11])
	fmt.Println(totients[87109])

	minRatio := float64(10000000)
	minN := 2
	for n := 2; n < Max; n++ {
		if LargestPermutation(n) == LargestPermutation(totients[n]) {
			ratio := float64(n) / float64(totients[n])
			if ratio < minRatio {
				minRatio = ratio
				minN = n
			}
		}
	}
	fmt.Println(minN)
}
