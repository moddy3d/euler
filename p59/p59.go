// XOR decryption
// https://projecteuler.net/problem=59

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	Letters = []rune{
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
)

func Sum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}

func main() {

	data, err := ioutil.ReadFile("cipher.txt")
	if err != nil {
		panic(err)
	}
	stringData := string(data)
	letters := strings.Split(strings.Trim(stringData, "\n"), ",")
	encrypted := []int{}
	for _, letter := range letters {
		i, err := strconv.Atoi(letter)
		if err != nil {
			panic(err)
		}
		encrypted = append(encrypted, i)
	}

	for a := 0; a < len(Letters); a++ {
		for b := 0; b < len(Letters); b++ {
			for c := 0; c < len(Letters); c++ {
				decrypted := []int{}
				for i := 0; i < len(encrypted)-2; i += 3 {
					part := encrypted[i : i+3]
					decryptPart := []int{}
					decryptPart = append(decryptPart, part[0]^int(Letters[a]))
					decryptPart = append(decryptPart, part[1]^int(Letters[b]))
					decryptPart = append(decryptPart, part[2]^int(Letters[c]))
					decrypted = append(decrypted, decryptPart...)
				}
				decrypted = append(decrypted, encrypted[len(encrypted)-1]^int(Letters[a]))

				runes := []rune{}
				for _, i := range decrypted {
					runes = append(runes, rune(i))
				}

				s := string(runes)

				words := strings.Split(s, " ")
				count := 0
				for _, word := range words {
					if word == "and" {
						count++
					}
				}
				if count > 0 {
					fmt.Println(s)
					fmt.Println(len(encrypted), len(decrypted))
					fmt.Println(Sum(decrypted))
					panic("stop")
				}
			}
		}
	}
}
