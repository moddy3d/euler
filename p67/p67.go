// Maximum path sum II
// https://projecteuler.net/problem=67

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	sampleData = [][]int{
		{3},
		{7, 4},
		{2, 4, 6},
		{8, 5, 9, 3},
	}
)

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func compute(data [][]int) int {
	for row := len(data) - 1; row > 0; row-- {
		cols := data[row]
		for col := 0; col < len(cols)-1; col++ {
			max := Max(cols[col], cols[col+1])
			data[row-1][col] += max
		}
	}

	return data[0][0]
}

func main() {

	data, err := ioutil.ReadFile("triangle.txt")
	if err != nil {
		panic(err)
	}
	stringData := string(data)
	rows := strings.Split(stringData, "\n")
	problemData := [][]int{}
	for _, row := range rows {
		ints := []int{}
		cols := strings.Split(row, " ")
		for _, col := range cols {
			i, err := strconv.Atoi(col)
			if err != nil {
				continue
			}
			ints = append(ints, i)
		}
		if len(ints) > 0 {
			problemData = append(problemData, ints)
		}
	}

	fmt.Println("Sample Data Solution:", compute(sampleData))
	fmt.Println("Problem Data Solution:", compute(problemData))
}
