// Magic 5-gon ring
// https://projecteuler.net/problem=68

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func Permutations(input []int, length int) [][]int {
	perms := [][]int{}

	if len(input) == 0 || length <= 0 {
		return perms
	}

	for i, val := range input {
		others := []int{}
		for j, v := range input {
			if i != j {
				others = append(others, v)
			}
		}

		if length == 1 {
			perms = append(perms, []int{val})
		} else {
			for _, p := range Permutations(others, length-1) {
				perm := append([]int{val}, p...)
				perms = append(perms, perm)
			}
		}
	}

	return perms
}

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func Min(numbers []int) int {
	min := numbers[0]
	for _, val := range numbers {
		if val < min {
			min = val
		}
	}
	return min
}

type Fivegon struct {
	data [][]int
}

func NewFivegon(values []int) *Fivegon {
	f := Fivegon{
		data: [][]int{},
	}

	for i := 0; i < len(values)/2; i++ {
		if i == len(values)/2-1 {
			f.data = append(f.data, []int{values[i+5], values[i], values[0]})
		} else {
			f.data = append(f.data, []int{values[i+5], values[i], values[i+1]})
		}
	}

	return &f
}

func (f *Fivegon) Digits() []int {

	fmt.Println(f.data)

	minIndex := 0
	min := f.data[0][0]
	for i, triplet := range f.data {
		key := triplet[0]
		if min > key {
			min = key
			minIndex = i
		}
	}

	fmt.Println(minIndex, min)

	digits := []int{}
	for _, triplet := range f.data[minIndex:] {
		fmt.Println(triplet)
		digits = append(digits, triplet...)
	}

	for _, triplet := range f.data[:minIndex] {
		fmt.Println(triplet)
		digits = append(digits, triplet...)
	}

	return digits
}

func (f *Fivegon) String() string {
	digits := f.Digits()
	runes := []rune{}
	for _, digit := range digits {
		runes = append(runes, []rune(strconv.Itoa(digit))...)
	}
	return string(runes)
}

func (f *Fivegon) Number() int {
	num, err := strconv.Atoi(f.String())
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	perms := Permutations([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 10)
	fmt.Println(len(perms))

	fivegons := []*Fivegon{}

	for _, perm := range perms {
		match := false
		matchSum := Sum([]int{perm[0], perm[1], perm[5]})
		for i := 1; i < len(perm)/2; i++ {
			sum := 0
			if i == len(perm)/2-1 {
				sum = Sum([]int{perm[i], perm[0], perm[i+5]})
				if sum == matchSum {
					match = true
				}
			} else {
				sum = Sum([]int{perm[i], perm[i+1], perm[i+5]})
				if sum != matchSum {
					break
				}
			}
		}

		if match {
			fivegons = append(fivegons, NewFivegon(perm))
		}
	}

	fmt.Println(fivegons[0].Number())

	numbers := []int{}
	for _, f := range fivegons {
		if len([]rune(f.String())) > 16 {
			continue
		}
		numbers = append(numbers, f.Number())
	}
	sort.Ints(numbers)

	fmt.Println(numbers)
}
