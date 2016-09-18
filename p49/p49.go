// Prime permutations
// https://projecteuler.net/problem=49

package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

var (
	PrimePermCache = map[int]*[]int{}
)

func Permutations(input []rune) [][]rune {
	perms := [][]rune{}

	for i, val := range input {
		others := []rune{}
		for j, v := range input {
			if j != i {
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

func PrimePerms(num int, primes []bool) []int {

	cache, exists := PrimePermCache[num]
	if exists {
		return *cache
	}

	str := strconv.Itoa(num)
	perms := Permutations([]rune(str))
	intMap := map[int]bool{}
	for _, p := range perms {
		n, err := strconv.Atoi(string(p))
		if err != nil {
			panic(err)
		}
		intMap[n] = true
	}

	ints := []int{}
	for key, _ := range intMap {
		if primes[key] == true && key > 1000 {
			ints = append(ints, key)
		}
	}

	// Sort it
	sort.Ints(ints)

	// Cache it
	for _, i := range ints {
		PrimePermCache[i] = &ints
	}

	return ints
}

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
	primes := PrimesUnder(10000)
	for i := 1000; i < 10000; i++ {
		if !primes[i] {
			continue
		}
		primePerms := PrimePerms(i, primes)
		for a := 0; a < len(primePerms); a++ {
			for b := a + 1; b < len(primePerms); b++ {
				for c := b + 1; c < len(primePerms); c++ {
					if primePerms[c]-primePerms[b] == primePerms[b]-primePerms[a] {
						fmt.Println(primePerms[a], primePerms[b], primePerms[c])
					}
				}
			}
		}
	}
}
