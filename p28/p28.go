// https://projecteuler.net/problem=28

package main

import (
	"fmt"
)

func main() {
	edgeLimit := 1001
	n := 1
	ring := 1
	sum := 1
	for {
		interval := ring * 2
		for i := 0; i < 4; i++ {
			n += interval
			sum += n
		}

		edge := interval + 1
		if edge == edgeLimit {
			break
		}

		ring += 1
	}
	fmt.Println(sum)
}
