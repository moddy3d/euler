package main

import "fmt"

func multiplicativeBinomialCoeff() {
	var (
		n     = 40
		k     = 20
		total = 1
	)

	for i := 0; i < k; i++ {
		total *= n - i
		total /= i + 1
	}
	fmt.Println(total)
}

func recurse(x, y int, cache *[20][20]int) int {
	if x == 0 || y == 0 {
		return 1
	}

	if cache[x-1][y-1] != 0 {
		return cache[x-1][y-1]
	} else {
		cache[x-1][y-1] = recurse(x, y-1, cache) + recurse(x-1, y, cache)
		return cache[x-1][y-1]
	}
}

func recursiveWithCache() {
	cache := &[20][20]int{}
	fmt.Println(recurse(20, 20, cache))
}

func main() {
	multiplicativeBinomialCoeff()
	recursiveWithCache()
}
