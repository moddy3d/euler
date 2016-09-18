// Prime digit replacements
// https://projecteuler.net/problem=51

package main

import (
	"fmt"
	"math"
	"strconv"
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
	primeMap := PrimesUnder(1000000)
	primes := []int{}
	for i, p := range primeMap {
		if p && i > 56003 {
			primes = append(primes, i)
		}
	}

	for _, p := range primes {
		runes := []rune(strconv.Itoa(p))
		for a := 0; a < len(runes); a++ {
			for b := a + 1; b < len(runes); b++ {
				for c := b + 1; c < len(runes); c++ {
					count := 0
					family := map[int]bool{}
					for i := 0; i <= 9; i++ {
						runes = []rune(strconv.Itoa(p))
						digit := []rune(strconv.Itoa(i))[0]
						runes[a] = digit
						runes[b] = digit
						runes[c] = digit
						num, err := strconv.Atoi(string(runes))
						if err != nil {
							panic(err)
						}
						if primeMap[num] && num >= p {
							family[num] = true
							count++
						}
					}
					_, exists := family[p]
					if count == 8 && exists {
						fmt.Println(p)
					}
				}
			}
		}
	}
}
