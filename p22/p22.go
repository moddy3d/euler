package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

func NameScore(name string) int {
	sum := 0
	for _, char := range name {
		sum += Score[string(char)]
	}
	return sum
}

func main() {
	data, err := ioutil.ReadFile("names.txt")
	if err != nil {
		panic(err)
	}
	stringData := string(data)
	strippedString := strings.Replace(stringData, "\"", "", -1)
	splitStrings := strings.Split(strippedString, ",")
	sort.Strings(splitStrings)

	sum := 0
	for i, name := range splitStrings {
		sum += (i + 1) * NameScore(name)
	}

	fmt.Println(sum)
}
