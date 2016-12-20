package main

import (
	"fmt"
	"time"
	"strings"
	"io/ioutil"
)

type cond struct {
	low int
	high int
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 15 GO!\n")

	var condition []cond
	var validIP int
	var first bool = true
	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")

	for _,e := range data {
		low := 0
		high := 0
		fmt.Sscanf(e, "%d-%d\n", &low, &high)
		condition = append(condition, cond{low, high})
	}

	for i:=0; i < 4294967295 ; i++ {
		valid := true
		for _, e := range condition {
			if i < e.low || i > e.high {
			} else{
				valid = false
				break
			}
		}
		if valid && first {
			fmt.Printf("Q1: Lowest IP is: %d\n", i)
			validIP++
			first = false
		} else if valid {
			validIP++
		}
	}
	
	fmt.Printf("Q2: Number of valid IP is: %d\n", validIP)

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
