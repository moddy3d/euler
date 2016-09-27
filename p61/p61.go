// Cyclical figurate numbers
// https://projecteuler.net/problem=61

package main

import (
	"fmt"
	"strconv"
)

var (
	Min int = 1000
	Max int = 9999
)

func SetBetween(f func(int) int, min, max int) []int {
	set := []int{}
	n := 0
	for {
		res := f(n)
		if res >= min && res <= max {
			set = append(set, res)
		} else if res > max {
			break
		}
		n++
	}
	return set
}

func MapBetween(f func(int) int, min, max int) map[string][]int {
	m := map[string][]int{}
	set := SetBetween(f, min, max)
	for _, n := range set {
		key := string([]rune(strconv.Itoa(n))[:2])
		s, exists := m[key]
		if !exists {
			s = []int{}
		}
		s = append(s, n)
		m[key] = s
	}
	return m
}

func Triangle(n int) int {
	return n * (n + 1) / 2
}

func Square(n int) int {
	return n * n
}

func Pentagonal(n int) int {
	return n * (3*n - 1) / 2
}

func Hexagonal(n int) int {
	return n * (2*n - 1)
}

func Heptagonal(n int) int {
	return n * (5*n - 3) / 2
}

func Octagonal(n int) int {
	return n * (3*n - 2)
}

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

func Find(n int, maps []map[string][]int, accum []int, matches *[][]int) {
	accum = append(accum, n)
	if len(maps) == 0 {
		if string([]rune(strconv.Itoa(accum[0]))[:2]) == string([]rune(strconv.Itoa(accum[len(accum)-1]))[2:]) {
			*matches = append(*matches, accum)
		}
		return
	}

	key := string([]rune(strconv.Itoa(n))[2:])
	m := maps[0]
	set, exists := m[key]
	if !exists {
		return
	}
	for _, n := range set {
		Find(n, maps[1:], accum, matches)
	}
}

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {
	sets := [][]int{
		SetBetween(Triangle, Min, Max),
		SetBetween(Square, Min, Max),
		SetBetween(Pentagonal, Min, Max),
		SetBetween(Hexagonal, Min, Max),
		SetBetween(Heptagonal, Min, Max),
		SetBetween(Octagonal, Min, Max),
	}

	maps := []map[string][]int{
		MapBetween(Triangle, Min, Max),
		MapBetween(Square, Min, Max),
		MapBetween(Pentagonal, Min, Max),
		MapBetween(Hexagonal, Min, Max),
		MapBetween(Heptagonal, Min, Max),
		MapBetween(Octagonal, Min, Max),
	}

	indices := []int{}
	for i, _ := range maps {
		indices = append(indices, i)
	}

	perms := Permutations(indices)
	matches := [][]int{}
	for _, perm := range perms {
		set := sets[perm[0]]

		orderedMaps := []map[string][]int{}
		for _, i := range perm[1:] {
			orderedMaps = append(orderedMaps, maps[i])
		}

		for _, n := range set {
			Find(n, orderedMaps, []int{}, &matches)
		}
	}

	sums := []int{}
	for _, set := range matches {
		sums = append(sums, Sum(set))
	}

	fmt.Println(sums)
}
