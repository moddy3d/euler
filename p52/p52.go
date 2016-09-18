// Permuted multiples
// https://projecteuler.net/problem=52

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func Sorted(num int) int {
	str := strconv.Itoa(num)
	ints := []int{}
	for _, r := range []rune(str) {
		i, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		ints = append(ints, i)
	}
	sort.Ints(ints)

	runes := []rune{}
	for _, i := range ints {
		runes = append(runes, []rune(strconv.Itoa(i))[0])
	}

	sortedInt, err := strconv.Atoi(string(runes))
	if err != nil {
		panic(err)
	}

	return sortedInt
}

func main() {
	n := 2
	for {
		count := 0
		s := Sorted(n)
		for f := 2; f <= 6; f++ {
			if Sorted(n*f) != s {
				break
			}
			count = f
		}
		if count == 6 {
			fmt.Println(n)
			break
		}
		n++
	}
}
