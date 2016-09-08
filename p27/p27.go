// https://projecteuler.net/problem=27

package main

import (
	"fmt"
	"math"
)

func Quadratic(a, b, n int) int {
	return int(math.Pow(float64(n), 2.0)) + a*n + b
}

func PrimeMapUnder(limit int) map[int]bool {
	numbers := map[int]bool{}
	for i := 2; i < limit; i++ {
		numbers[i] = true
	}

	for i := 2; i < limit; i++ {
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

func PrimesUnder(primeMap map[int]bool) []int {
	primes := []int{}
	for key, value := range primeMap {
		if value == true {
			primes = append(primes, key)
		}
	}
	return primes
}

func ExtendNegativePart(positives []int) []int {
	numbers := []int{}
	for _, n := range positives {
		numbers = append(numbers, n)
		numbers = append([]int{-n}, numbers...)
	}
	return numbers
}

func main() {
	quadraticPrimeMap := PrimeMapUnder(Quadratic(1000, 1000, 79))

	primeMap := PrimeMapUnder(1001)
	absolutePrimes := ExtendNegativePart(PrimesUnder(primeMap))
	highestConsecutive := 0
	product := 0
	for _, a := range absolutePrimes {
		for _, b := range absolutePrimes {
			n := 0
			for {
				quadratic := Quadratic(a, b, n)
				if quadraticPrimeMap[quadratic] == false {
					break
				}
				n += 1
			}
			if n > highestConsecutive {
				highestConsecutive = n
				product = a * b
				pa = a
				pb = b
			}

		}
	}
	fmt.Println(highestConsecutive, product)
}
