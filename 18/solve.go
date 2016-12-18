package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

var room [400000][]string

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 18 GO!\n")

	var left, center, right string
	var safe int

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")
	for _, e := range data[0] {
		room[0] = append(room[0], string(e))
	}

	safe = strings.Count(strings.Join(room[0], ""), ".")

	for i := 1; i < len(room); i++ {
		var row []string
		for j := 0; j < len(room[i-1]); j++ {
			if j == 0 {
				left = "."
				center = room[i-1][j]
				right = room[i-1][j+1]
			} else if j == len(room[i-1])-1 {
				left = room[i-1][j-1]
				center = room[i-1][j]
				right = "."
			} else {
				left = room[i-1][j-1]
				center = room[i-1][j]
				right = room[i-1][j+1]
			}
			if left == "^" && center == "^" && right == "." {
				row = append(row, "^")
			} else if left == "." && center == "^" && right == "^" {
				row = append(row, "^")
			} else if left == "^" && center == "." && right == "." {
				row = append(row, "^")
			} else if left == "." && center == "." && right == "^" {
				row = append(row, "^")
			} else {
				row = append(row, ".")
			}
		}
		room[i] = row
		safe += strings.Count(strings.Join(room[i], ""), ".")
		if i == 39 {
			fmt.Printf("Q1: Number of safe tiles is %d\n", safe)
		}
	}

	fmt.Printf("Q2: Number of safe tiles is %d\n", safe)

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
