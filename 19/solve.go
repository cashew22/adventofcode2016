package main

import (
	"fmt"
	"time"
)

func safePlace(n, k int) int {
	if n == 1 {
		return 0
	} else {
		return (safePlace(n-1, k) + k) % n
	}
}

func safePlaceDay2(n int) int {
	i := 1
	for i*3 < n {
		i *= 3
	}
	return n - i
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 16 GO!\n")

	var input int = 3004953
	luckyElf := safePlace(input, 2) + 1
	fmt.Printf("Q1: The elf with all the gift is the one in position: %d\n", luckyElf)
	luckyElf = safePlaceDay2(input)
	fmt.Printf("Q2: The elf with all the gift is the one in position: %d\n", luckyElf)

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
