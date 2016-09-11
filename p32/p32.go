// https://projecteuler.net/problem=32

package main

import (
	"fmt"
	"math"
)

var (
	digits = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
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

func SliceToInt(slice []int) int {
	sum := 0
	for i, val := range slice {
		sum += val * int(math.Pow(10, float64(len(slice)-i-1)))
	}
	return sum
}

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {
	productMap := map[int]bool{}
	perms := Permutations(digits)

	for _, perm := range perms {
		a := SliceToInt(perm[:1])
		b := SliceToInt(perm[1:5])
		p := SliceToInt(perm[5:])
		if a*b == p {
			productMap[p] = true
		}
	}

	for _, perm := range perms {
		a := SliceToInt(perm[:2])
		b := SliceToInt(perm[2:5])
		p := SliceToInt(perm[5:])
		if a*b == p {
			productMap[p] = true
		}
	}

	products := []int{}
	for key, _ := range productMap {
		products = append(products, key)
	}

	fmt.Println(Sum(products))
}
