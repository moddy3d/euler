// Distinct primes factors
// https://projecteuler.net/problem=47

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

func DistinctPrimeFactors(numbers []int, primes *[]bool) bool {
	for _, number := range numbers {
		primeFactors := map[int]bool{}
		tmp := float64(number)
		for i := 2; i < number; i++ {
			if (*primes)[i] {
				if math.Mod(tmp, float64(i)) == 0.0 {
					tmp /= float64(i)
					primeFactors[i] = true
					i = 1
				}
			}
		}
		if len(primeFactors) != len(numbers) {
			return false
		}
	}
	return true
}

func main() {
	primes := PrimesUnder(1000000)
	start := 645
	for {
		numbers := []int{start, start + 1, start + 2, start + 3}
		if DistinctPrimeFactors(numbers, &primes) {
			fmt.Println(numbers[0])
			break
		}
		start += 1
	}
}
