// Champernowne's constant
// https://projecteuler.net/problem=40

package main

import (
	"fmt"
	"strconv"
)

func main() {
	champ := []rune{}
	for n := 1; n <= 1000000; n++ {
		champ = append(champ, []rune(strconv.Itoa(n))...)
	}

	indices := []int{0, 9, 99, 999, 9999, 99999, 999999}
	product := 1
	for _, i := range indices {
		n, _ := strconv.Atoi(string(champ[i]))
		product *= n
	}
	fmt.Println(product)
}
