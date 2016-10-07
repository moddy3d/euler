// Counting fractions
// https://projecteuler.net/problem=72

package main

import "fmt"

func main() {
	var (
		Max int = 1000000
	)

	accum := make([]int, Max+1, Max+1)
	for i, _ := range accum {
		accum[i] = i
	}

	var count = 0
	for d := 2; d <= Max; d++ {
		if accum[d] == d {
			for n := d; n <= Max; n += d {
				accum[n] = accum[n] - (accum[n] / d)
			}
		}
		count += accum[d]
	}
	fmt.Println(count)

}
