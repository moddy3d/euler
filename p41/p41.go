// Pandigital prime
// https://projecteuler.net/problem=41

package main

import (
	"fmt"
	"math"
	"strconv"
)

func Permutations(input []rune) [][]rune {
	perms := [][]rune{}

	for _, val := range input {
		others := []rune{}
		for _, v := range input {
			if v != val {
				others = append(others, v)
			}
		}

		if len(others) > 0 {
			for _, p := range Permutations(others) {
				perm := append([]rune{val}, p...)
				perms = append(perms, perm)
			}
		} else {
			perms = append(perms, input)
		}
	}

	return perms
}

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

func main() {
	primes := PrimesUnder(7654322)

	perms := Permutations([]rune{'7', '6', '5', '4', '3', '2', '1'})
	permInts := []int{}
	for _, p := range perms {
		n, _ := strconv.Atoi(string(p))
		permInts = append(permInts, n)
	}

	greatest := 0
	for _, p := range permInts {
		if primes[p] == true {
			if p > greatest {
				greatest = p
			}
		}
	}
	fmt.Println(greatest)
}
