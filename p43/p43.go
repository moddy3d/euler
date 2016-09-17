// Sub-string divisibility
// https://projecteuler.net/problem=43

package main

import (
	"fmt"
	"math"
	"strconv"
)

var (
	Divisors = []int{1, 2, 3, 5, 7, 11, 13, 17}
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

func SubStringDivisible(seq []rune) bool {
	for s := 1; s < len(seq)-2; s++ {
		sub := seq[s : s+3]
		num, _ := strconv.Atoi(string(sub))
		if math.Mod(float64(num), float64(Divisors[s])) != 0.0 {
			return false
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
	// Build permutations
	perms := Permutations([]rune{'9', '8', '7', '6', '5', '4', '3', '2', '1', '0'})
	fmt.Println("Got perms")

	// Find divisibles
	divisibles := [][]rune{}
	for _, perm := range perms {
		if perm[0] != '0' && SubStringDivisible(perm) {
			divisibles = append(divisibles, perm)
		}
	}

	// Convert to integers
	ints := []int{}
	for _, d := range divisibles {
		n, _ := strconv.Atoi(string(d))
		ints = append(ints, n)
	}

	// Sum it
	fmt.Println(Sum(ints))
}
