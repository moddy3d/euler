// Coded triangle numbers
// https://projecteuler.net/problem=42

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	Score = map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12,
		"M": 13,
		"N": 14,
		"O": 15,
		"P": 16,
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	}
)

func WordScore(name string) int {
	sum := 0
	for _, char := range name {
		sum += Score[string(char)]
	}
	return sum
}

func TriangleNumbersUnder(n int) map[int]bool {
	tris := map[int]bool{}
	for i := 0; i <= n; i++ {
		tris[i] = false
	}

	i := 1
	for {
		t := int(float64(i) / 2.0 * (float64(i) + 1.0))
		if t > n {
			break
		}
		tris[t] = true
		i += 1
	}

	return tris
}

func main() {
	data, err := ioutil.ReadFile("words.txt")
	if err != nil {
		panic(err)
	}
	stringData := string(data)
	strippedString := strings.Replace(stringData, "\"", "", -1)
	words := strings.Split(strippedString, ",")

	tris := TriangleNumbersUnder(192)

	triWords := 0
	for _, word := range words {
		if tris[WordScore(word)] == true {
			triWords += 1
		}
	}

	fmt.Println(triWords)
}
