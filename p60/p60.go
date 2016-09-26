// Prime pair sets
// https://projecteuler.net/problem=60

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

func PrimePair(a, b int, primeMap []bool) bool {
	ab := a*Padding(b) + b
	if !primeMap[ab] {
		return false
	}

	ba := b*Padding(a) + a
	if !primeMap[ba] {
		return false
	}

	return true
}

func Padding(n int) int {
	var p int = 1
	for p < n {
		p *= 10
	}
	return p
}

func PrimePairSet(set []int, primeMap []bool) bool {
	for a := 0; a < len(set); a++ {
		for b := a + 1; b < len(set); b++ {
			if !PrimePair(set[a], set[b], primeMap) {
				return false
			}
		}
	}
	return true
}

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {
	primeMap := PrimesUnder(300000000)
	primes := []int{}
	for i, isPrime := range primeMap {
		if isPrime && i < 10000 {
			primes = append(primes, i)
		}
	}

	twoSets := [][]int{}
	for a := 0; a < len(primes); a++ {
		for b := a + 1; b < len(primes); b++ {
			set := []int{
				primes[a],
				primes[b],
			}
			if PrimePairSet(set, primeMap) {
				twoSets = append(twoSets, set)
			}
		}
	}

	fmt.Println("Two sets done")

	threeSets := [][]int{}
	for a := 0; a < len(twoSets); a++ {
		for b := 0; b < len(primes); b++ {
			set := []int{
				twoSets[a][0],
				twoSets[a][1],
				primes[b],
			}
			if PrimePairSet(set, primeMap) {
				threeSets = append(threeSets, set)
			}
		}
	}

	fmt.Println("Three sets done")

	fiveSets := [][]int{}
	for a := 0; a < len(threeSets); a++ {
		for b := 0; b < len(twoSets); b++ {
			set := []int{
				threeSets[a][0],
				threeSets[a][1],
				threeSets[a][2],
				twoSets[b][0],
				twoSets[b][1],
			}

			if PrimePairSet(set, primeMap) {
				fiveSets = append(fiveSets, set)
			}

		}
	}

	fmt.Println("Five sets done")

	lowestSum := 99999999999999999
	for _, set := range fiveSets {
		sum := Sum(set)
		if sum < lowestSum {
			lowestSum = sum
		}
	}
	fmt.Println(lowestSum)
}
