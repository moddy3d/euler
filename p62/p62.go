// Cubic permutations
// https://projecteuler.net/problem=62

package main

import (
	"fmt"
	"math"
)

type CubeCount struct {
	SmallestCube int
	Count        int
}

func Cube(n int) int {
	return n * n * n
}

func CubesUnder(limit int) []bool {
	cubes := make([]bool, limit+1, limit+1)

	n := 0
	for {
		cube := Cube(n)
		if cube > limit {
			break
		}
		cubes[cube] = true
		n++
	}

	return cubes
}

func LargestPermutation(n int) int {
	digits := make([]int, 10)
	perm := 0

	k := n
	for k > 0 {
		digits[int(math.Mod(float64(k), 10.0))]++
		k /= 10
	}

	for i := 9; i >= 0; i-- {
		for j := 0; j < digits[i]; j++ {
			perm = perm*10 + i
		}
	}

	return perm
}

func main() {
	counter := map[int]*CubeCount{}
	n := 0
	for {
		cube := Cube(n)
		perm := LargestPermutation(cube)

		count, exists := counter[perm]
		if !exists {
			count = &CubeCount{cube, 1}
			counter[perm] = count
		} else {
			count.Count++
			if cube < count.SmallestCube {
				count.SmallestCube = cube
			}
		}

		if count.Count == 5 {
			fmt.Println(n, count.SmallestCube)
			break
		}

		n++
	}

}
