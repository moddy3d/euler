// https://projecteuler.net/problem=39

package main

import (
	"fmt"
	"math"
)

func RightTriangles(perimeter int) int {
	tris := 0
	for a := 1; a <= perimeter/2; a++ {
		for b := 1; b <= perimeter/2; b++ {
			cSquared := math.Pow(float64(a), 2.0) + math.Pow(float64(b), 2.0)
			cFloat := math.Sqrt(cSquared)
			if cFloat == float64(int64(cFloat)) {
				c := int(cFloat)
				if a+b+c == perimeter {
					tris += 1
				}
			}
		}
	}
	return tris / 2
}

func main() {
	greatest := 0
	greatestP := 0
	for p := 1; p <= 1000; p++ {
		tris := RightTriangles(p)
		if tris > greatest {
			greatest = tris
			greatestP = p
		}
	}
	fmt.Println(greatestP)
}
