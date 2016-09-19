// Poker hands
// https://projecteuler.net/problem=54

package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

var (
	CardValue = map[rune]int{
		'2': 2,
		'3': 3,
		'4': 4,
		'5': 5,
		'6': 6,
		'7': 7,
		'8': 8,
		'9': 9,
		'T': 10,
		'J': 11,
		'Q': 12,
		'K': 13,
		'A': 14,
	}

	SuitValue = map[rune]int{
		'D': 1,
		'C': 2,
		'H': 3,
		'S': 4,
	}
)

type Card struct {
	Value int
	Suit  int
}

func Cards(cardStrings []string) []*Card {
	cards := []*Card{}
	for _, str := range cardStrings {
		runes := []rune(str)
		cards = append(cards, &Card{CardValue[runes[0]], SuitValue[runes[1]]})
	}
	return cards
}

func RoyalFlush(cards []*Card) bool {
	suits := map[int]bool{}
	for _, card := range cards {
		suits[card.Suit] = true
	}

	if len(suits) > 1 {
		return false
	}

	for _, card := range cards {
		if card.Value < 10 {
			return false
		}
	}

	return true
}

func StraightFlush(cards []*Card) (int, bool) {
	suits := map[int]bool{}
	for _, card := range cards {
		suits[card.Suit] = true
	}

	if len(suits) > 1 {
		return 0, false
	}

	values := []int{}
	for _, card := range cards {
		values = append(values, card.Value)
	}
	sort.Ints(values)
	for i := 0; i < len(values)-1; i++ {
		if values[i+1]-values[i] != 1 {
			return 0, false
		}
	}

	return values[4], true
}

func FourOfAKind(cards []*Card) (int, bool) {
	values := map[int]int{}
	for _, card := range cards {
		values[card.Value]++
	}

	for value, count := range values {
		if count == 4 {
			return value, true
		}
	}

	return 0, false
}

func FullHouse(cards []*Card) (int, bool) {
	values := map[int]int{}
	for _, card := range cards {
		values[card.Value]++
	}

	highestA := 0
	matchA := false
	matchB := false
	for value, count := range values {
		if count == 3 {
			matchA = true
			highestA = value
		} else if count == 2 {
			matchB = true
		}
	}

	if matchA && matchB {
		return highestA, true
	}

	return 0, false
}

func Flush(cards []*Card) (int, bool) {
	suits := map[int]bool{}
	highest := 0
	for _, card := range cards {
		suits[card.Suit] = true
		if card.Value > highest {
			highest = card.Value
		}
	}

	if len(suits) == 1 {
		return highest, true
	}

	return 0, false
}

func Straight(cards []*Card) (int, bool) {
	values := []int{}
	for _, card := range cards {
		values = append(values, card.Value)
	}
	sort.Ints(values)
	for i := 0; i < len(values)-1; i++ {
		if values[i+1]-values[i] != 1 {
			return 0, false
		}
	}

	return values[4], true
}

func ThreeOfAKind(cards []*Card) (int, bool) {
	values := map[int]int{}
	for _, card := range cards {
		values[card.Value]++
	}

	highest := 0
	match := false
	for value, count := range values {
		if count == 3 {
			match = true
			highest = value
		}
	}

	if match {
		return highest, true
	}

	return 0, false
}

func TwoPairs(cards []*Card) (int, int, bool) {
	values := map[int]int{}
	for _, card := range cards {
		values[card.Value]++
	}

	var (
		highestA int
		highestB int
		matchA   bool
		matchB   bool
	)
	for value, count := range values {
		if count == 2 {
			if !matchA {
				matchA = true
				highestA = value
			} else {
				matchB = true
				highestB = value
			}
		}
	}

	if matchA && matchB {
		if highestA > highestB {
			return highestA, highestB, true
		} else if highestB > highestA {
			return highestB, highestA, true
		}
	}

	return 0, 0, false
}

func OnePair(cards []*Card) (int, bool) {
	values := map[int]int{}
	for _, card := range cards {
		values[card.Value]++
	}

	var (
		highest int
		match   bool
	)
	for value, count := range values {
		if count == 2 {
			match = true
			highest = value
		}
	}

	if match {
		return highest, true
	}

	return 0, false
}

func OrderedValues(cards []*Card) []int {
	values := []int{}
	for _, card := range cards {
		values = append(values, card.Value)
	}
	sort.Ints(values)
	return values
}

func main() {
	data, err := ioutil.ReadFile("poker.txt")
	if err != nil {
		panic(err)
	}
	stringData := string(data)
	rounds := strings.Split(stringData, "\n")

	score1 := 0
	score2 := 0
	for _, round := range rounds {
		cardStrings := strings.Split(round, " ")
		if len(cardStrings) != 10 {
			continue
		}
		cards1 := Cards(cardStrings[:5])
		cards2 := Cards(cardStrings[5:])

		// Royal Flush
		match1 := RoyalFlush(cards1)
		match2 := RoyalFlush(cards2)
		switch {
		case match1 && !match2:
			score1++
			continue
		case match2 && !match1:
			score2++
			continue
		}

		// Straight Flush
		highest1, match1 := StraightFlush(cards1)
		highest2, match2 := StraightFlush(cards2)
		switch {
		case match1 && !match2:
			score1++
			continue
		case match2 && !match1:
			score2++
			continue
		case match1 && match2:
			if highest1 > highest2 {
				score1++
				continue
			} else if highest2 > highest1 {
				score2++
				continue
			}
		}

		// Four of a kind
		highest1, match1 = FourOfAKind(cards1)
		highest2, match2 = FourOfAKind(cards2)
		switch {
		case match1 && !match2:
			score1++
			continue
		case match2 && !match1:
			score2++
			continue
		case match1 && match2:
			if highest1 > highest2 {
				score1++
				continue
			} else if highest2 > highest1 {
				score2++
				continue
			}
		}

		// FullHouse
		highest1, match1 = FullHouse(cards1)
		highest2, match2 = FullHouse(cards2)
		switch {
		case match1 && !match2:
			score1++
			continue
		case match2 && !match1:
			score2++
			continue
		case match1 && match2:
			if highest1 > highest2 {
				score1++
				continue
			} else if highest2 > highest1 {
				score2++
				continue
			}
		}

		// Flush
		highest1, match1 = Flush(cards1)
		highest2, match2 = Flush(cards2)
		switch {
		case match1 && !match2:
			score1++
			continue
		case match2 && !match1:
			score2++
			continue
		case match1 && match2:
			if highest1 > highest2 {
				score1++
				continue
			} else if highest2 > highest1 {
				score2++
				continue
			}
		}

		// Straight
		highest1, match1 = Straight(cards1)
		highest2, match2 = Straight(cards2)
		switch {
		case match1 && !match2:
			score1++
			continue
		case match2 && !match1:
			score2++
			continue
		case match1 && match2:
			if highest1 > highest2 {
				score1++
				continue
			} else if highest2 > highest1 {
				score2++
				continue
			}
		}

		// Three of a kind
		highest1, match1 = ThreeOfAKind(cards1)
		highest2, match2 = ThreeOfAKind(cards2)
		switch {
		case match1 && !match2:
			score1++
			continue
		case match2 && !match1:
			score2++
			continue
		case match1 && match2:
			if highest1 > highest2 {
				score1++
				continue
			} else if highest2 > highest1 {
				score2++
				continue
			}
		}

		// Two Pairs
		highestA1, highestB1, match1 := TwoPairs(cards1)
		highestA2, highestB2, match2 := TwoPairs(cards2)
		switch {
		case match1 && !match2:
			score1++
			continue
		case match2 && !match1:
			score2++
			continue
		case match1 && match2:
			if highestA1 > highestA2 {
				score1++
				continue
			} else if highestA1 == highestA2 {
				if highestB1 > highestB2 {
					score1++
					continue
				} else if highestB2 > highestB1 {
					score2++
					continue
				}
			} else {
				score2++
				continue
			}
		}

		// One Pair
		highest1, match1 = OnePair(cards1)
		highest2, match2 = OnePair(cards2)
		switch {
		case match1 && !match2:
			score1++
			continue
		case match2 && !match1:
			score2++
			continue
		case match1 && match2:
			if highest1 > highest2 {
				score1++
				continue
			} else if highest2 > highest1 {
				score2++
				continue
			}
		}

		// Highest Card
		ordered1 := OrderedValues(cards1)
		ordered2 := OrderedValues(cards2)
		for i := 4; i >= 0; i++ {
			if ordered1[i] > ordered2[i] {
				score1++
				break
			} else if ordered2[i] > ordered1[i] {
				score2++
				break
			}
		}
	}
	fmt.Println(score1, score2)

}
