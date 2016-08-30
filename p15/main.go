package main

import "fmt"

var (
	n = 40
	k = 20
)

func main() {
	total := 1
	for i := 0; i < k; i++ {
		total *= n - i
		total /= i + 1
	}
	fmt.Println(total)
}
