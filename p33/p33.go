// https://projecteuler.net/problem=33

package main

import (
	"fmt"
	"math"
	"strconv"
)

type Fraction struct {
	Numerator   int
	Denominator int
}

func PopLeft(n int) (int, int) {
	num, err := strconv.Atoi(string(strconv.Itoa(n)[1:]))
	if err != nil {
		panic(err)
	}
	pop, err := strconv.Atoi(string(strconv.Itoa(n)[:1]))
	if err != nil {
		panic(err)
	}
	return num, pop
}

func PopRight(n int) (int, int) {
	runes := strconv.Itoa(n)
	num, err := strconv.Atoi(string(runes[:len(runes)-1]))
	if err != nil {
		panic(err)
	}
	pop, err := strconv.Atoi(string(runes[len(runes)-1:]))
	if err != nil {
		panic(err)
	}
	return num, pop
}

func Check(f1, f2 *Fraction) bool {
	if f2.Numerator == 0 || f2.Denominator == 0 {
		return false
	}
	return float64(f1.Numerator)/float64(f2.Numerator) == float64(f1.Denominator)/float64(f2.Denominator)
}

func Cancel(f *Fraction) bool {
	var (
		check bool
		a     int
		ap    int
		b     int
		bp    int
	)

	if f.Numerator == f.Denominator {
		return false
	}

	if int(math.Mod(float64(f.Numerator), 10.0)) == 0 && int(math.Mod(float64(f.Denominator), 10.0)) == 0 {
		return false
	}

	a, ap = PopRight(f.Numerator)
	b, bp = PopLeft(f.Denominator)
	check = Check(f, &Fraction{a, b})
	if ap == bp && check {
		return true
	}

	return false
}

func main() {
	curious := []*Fraction{}
	for a := 10; a < 100; a++ {
		for b := 10; b < 100; b++ {
			f := &Fraction{a, b}
			if Cancel(f) {
				curious = append(curious, f)
			}
		}
	}

	for _, k := range curious {
		fmt.Println(*k)
	}

	numerator := 1
	denominator := 1
	for _, k := range curious {
		numerator *= k.Numerator
		denominator *= k.Denominator
	}
	fmt.Println(denominator / numerator)
}
