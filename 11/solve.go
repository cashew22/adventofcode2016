package main

import (
	"fmt"
	"time"
)

//THX INTERNET
func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 11 GO !\n")

	var (
		itemOnFloor2 = 2
		item         = 10
		itemDay2     = 14
		stepToRemove = 0
	)

	step := (2*(item-1) - 1) * 3
	stepDay2 := (2*(itemDay2-1) - 1) * 3
	stepToRemove = (2*(itemOnFloor2) - 1) + 1

	fmt.Printf("Q1: Minimal required steps = %d\n", step-stepToRemove)
	fmt.Printf("Q2: Minimal required steps = %d\n", stepDay2-stepToRemove)

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
