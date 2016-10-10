// Counting fractions in a range
// https://projecteuler.net/problem=73

package main

import "fmt"

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}

func main() {
	var (
		Max int = 12000
	)

	var count = 0
	for d := 4; d <= Max; d++ {
		for n := d/3 + 1; n < (d-1)/2+1; n++ {
			if GCD(n, d) == 1 {
				count++
			}
		}
	}
	fmt.Println(count)

}
