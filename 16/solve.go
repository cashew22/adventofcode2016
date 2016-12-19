package main

import (
	"fmt"
	"time"
)

func reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func turnBit(b []int) []int {
	for i := 0; i < len(b); i++ {
		if b[i] == 0 {
			b[i] = 1
		} else {
			b[i] = 0
		}
	}
	return b
}

func fill(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	b = turnBit(reverse(b))
	a = append(a, 0)
	return append(a, b...)
}

func checksum(a []int) []int {
	var b []int
	for i := 0; i < len(a)-1; i += 2 {
		if a[i] == a[i+1] {
			b = append(b, 1)
		} else {
			b = append(b, 0)
		}
	}
	return b
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 16 GO!\n")

	input := []int{0, 1, 0, 0, 0, 1, 0, 0, 0, 1, 0, 0, 1, 0, 1, 1, 1}
	lenght := 35651584

	for len(input) < lenght {
		input = fill(input)
	}

	input = checksum(input[:lenght])
	for len(input)%2 == 0 {
		input = checksum(input)
	}

	fmt.Printf("The checksum is: ")
	for _, e := range input {
		fmt.Printf("%d", e)
	}
	fmt.Printf("\n")

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
