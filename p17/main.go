package main

import (
	"fmt"
	"strconv"
)

var (
	numbers = map[int]string{
		0:    "",
		1:    "one",
		2:    "two",
		3:    "three",
		4:    "four",
		5:    "five",
		6:    "six",
		7:    "seven",
		8:    "eight",
		9:    "nine",
		10:   "ten",
		11:   "eleven",
		12:   "twelve",
		13:   "thirteen",
		14:   "fourteen",
		15:   "fifteen",
		16:   "sixteen",
		17:   "seventeen",
		18:   "eighteen",
		19:   "nineteen",
		20:   "twenty",
		30:   "thirty",
		40:   "forty",
		50:   "fifty",
		60:   "sixty",
		70:   "seventy",
		80:   "eighty",
		90:   "ninety",
		100:  "hundred",
		1000: "onethousand",
	}
	and   string = "and"
	start int    = 1
	end   int    = 1000
)

func lenForNum(i int) int {
	var (
		letters int
		word    string
		exists  bool
		tmp     int
		x       int
		xx      int
		x00     int
		err     error
	)

	// If our number is defined then return early
	word, exists = numbers[i]
	if exists {
		if i == 100 {
			return len("onehundred")
		}
		return len(word)
	}

	numStr := strconv.Itoa(i)
	l := len(string([]rune(numStr)))

	x, err = strconv.Atoi(fmt.Sprintf("%s", numStr[l-1:]))
	if err != nil {
		panic(err)
	}
	if l > 1 {
		tmp, err = strconv.Atoi(fmt.Sprintf("%s", numStr[l-2:l-1]))
		if err != nil {
			panic(err)
		}
		xx = tmp * 10
	}
	if l > 2 {
		tmp, err = strconv.Atoi(fmt.Sprintf("%s", numStr[l-3:l-2]))
		if err != nil {
			panic(err)
		}
		x00 = tmp
	}

	if xx+x < 20 {
		letters += len(numbers[xx+x])
	} else {
		letters += len(numbers[xx])
		letters += len(numbers[x])
	}

	if x00 > 0 {
		if xx+x > 0 {
			letters += len(and)
		}
		letters += len(numbers[x00])
		letters += len(numbers[100])
	}

	return letters
}

func main() {
	letters := 0
	for i := start; i <= end; i++ {
		length := lenForNum(i)
		letters += length
		fmt.Println(i, length)
	}
	fmt.Println(letters)
}
