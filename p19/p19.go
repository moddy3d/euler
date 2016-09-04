package main

import (
	"fmt"
	"math"
)

// 1 Jan 1901 fell on Tuesday

var (
	DaysInMonth = map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	LeapYearDaysInMonth = map[int]int{
		1:  31,
		2:  29,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	}
	StartYear = 1901
	EndYear   = 2000
)

func LeapYear(year int) bool {
	if math.Mod(float64(year), 100) == 0 {
		if math.Mod(float64(year), 400) == 0 {
			return true
		} else {
			return false
		}
	} else {
		if math.Mod(float64(year), 4) == 0 {
			return true
		} else {
			return false
		}
	}

}

func main() {
	var (
		count = 0
		day   = 2
	)
	for year := StartYear; year <= EndYear; year++ {
		var daysInMonth map[int]int
		if LeapYear(year) {
			daysInMonth = LeapYearDaysInMonth
		} else {
			daysInMonth = DaysInMonth
		}
		for month := 1; month <= 12; month++ {
			day += daysInMonth[month]
			if math.Mod(float64(day), 7) == 0 {
				count += 1
			}
		}
	}

	fmt.Println(count)
}
