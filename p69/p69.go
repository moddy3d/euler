// Totient maximum
// https://projecteuler.net/problem=69

package main

import "fmt"

func main() {
	totients := make([]int, 1000001, 1000001)
	used := make([]bool, 1000001, 1000001)
	for n := 0; n < 1000001; n++ {
		totients[n] = n
	}
	totients[0] = 1
	totients[1] = 1

	for n := 2; n < 1000001; n++ {
		factors := 0
		multiple := 1
		for {
			factor := n * multiple
			if factor < 1000001 {
				if !used[factor] {
					factors++
					used[factor] = true
				}
				if factor == 6 {
					fmt.Println(factors)
				}
				totients[factor] -= factors
			} else {
				break
			}
			multiple++
		}
	}
	fmt.Println(totients[2:11])

	maxRatio := float64(0.0)
	maxN := 0
	for n := 1; n < 1000000; n++ {
		ratio := float64(n) / float64(totients[n])
		if ratio > maxRatio {
			maxRatio = ratio
			maxN = n
		}
	}
	fmt.Println(maxN)
}
