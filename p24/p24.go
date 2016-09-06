// https://projecteuler.net/problem=24

package main

import (
	"fmt"
)

var (
	sampleInput  = []int{0, 1, 2}
	problemInput = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func Permutations(input []int) [][]int {
	perms := [][]int{}

	for _, val := range input {
		others := []int{}
		for _, v := range input {
			if v != val {
				others = append(others, v)
			}
		}

		if len(others) > 0 {
			for _, p := range Permutations(others) {
				perm := append([]int{val}, p...)
				perms = append(perms, perm)
			}
		} else {
			perms = append(perms, input)
		}
	}

	return perms
}

func main() {
	perms := Permutations(problemInput)
	fmt.Println(perms[999999])
}
