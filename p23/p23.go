package main

import (
	"fmt"
	"math"
)

type NumberType int

var (
	nonabundantLimit = 28123
)

const (
	Perfect NumberType = iota
	Abundant
	Deficient
	Unknown
)

// Returns the proper divisors of a number
func ProperDivisors(number int) *[]int {
	divisorMap := map[int]bool{}
	lim := number
	for i := 1; i < lim; i++ {
		if math.Mod(float64(number), float64(i)) == 0 {
			divisorMap[i] = true
			pair := number / i
			lim = pair
			divisorMap[pair] = true
		}
	}

	divisors := []int{}
	for k, _ := range divisorMap {
		if k != number {
			divisors = append(divisors, k)
		}
	}

	return &divisors
}

// Returns the sum of a slice of ints
func Sum(numbers *[]int) int {
	sum := 0
	for _, val := range *numbers {
		sum += val
	}
	return sum
}

// Returns the NumberType of a number
func Classify(number int) NumberType {
	sum := Sum(ProperDivisors(number))
	switch {
	case sum == number:
		return Perfect
	case sum > number:
		return Abundant
	case sum < number:
		return Deficient
	default:
		return Unknown
	}
}

func main() {
	abundantNumbers := []int{}
	for i := 12; i <= nonabundantLimit; i++ {
		if Classify(i) == Abundant {
			abundantNumbers = append(abundantNumbers, i)
		}
	}

	abundantSumMap := map[int]bool{}
	for _, a := range abundantNumbers {
		for _, b := range abundantNumbers {
			sum := a + b
			if sum > nonabundantLimit {
				break
			} else {
				abundantSumMap[sum] = true
			}
		}
	}

	solution := []int{}
	for i := 1; i <= nonabundantLimit; i++ {
		_, exists := abundantSumMap[i]
		if !exists {
			solution = append(solution, i)
		}
	}

	fmt.Println(Sum(&solution))
}
