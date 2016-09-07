// https://projecteuler.net/problem=26

package main

import (
	"fmt"
	"math"
)

func main() {
	numberWithLongestCycle := 0
	longestCycle := 0
	for i := 1000; i > 0; i-- {
		remainders := map[int]int{}
		dividend := 1
		cycle := 0
		for {
			remainders[dividend] = cycle
			dividend = dividend * 10
			dividend = int(math.Remainder(float64(dividend), float64(i)))
			cycle += 1

			c := remainders[dividend]
			if c != 0 || dividend == 0 {
				break
			}
		}

		c := remainders[dividend]
		if cycle-c > longestCycle {
			longestCycle = cycle
			numberWithLongestCycle = i
		}
	}
}
