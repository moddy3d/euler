// Spiral primes
// https://projecteuler.net/problem=58

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

func main() {
	primes := PrimesUnder(800000000)

	var (
		n          int = 1
		ring       int = 1
		diags      int = 1
		primeDiags int = 0
	)

	for {
		interval := ring * 2
		for i := 0; i < 4; i++ {
			n += interval
			diags += 1
			if primes[n] {
				primeDiags += 1
			}
		}

		edge := interval + 1
		ring += 1

		fmt.Printf("%d/%d\n", primeDiags, diags)

		if float64(primeDiags)/float64(diags) < 0.1 {
			fmt.Println(edge)
			break
		}
	}
}
