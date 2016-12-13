package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func day12Algo(data []string, day2 bool) int {
	var (
		reg    map[string]int
		index  = 0
		source string
		dest   string
	)
	reg = make(map[string]int)

	if day2 {
		reg["c"] = 1
	}

	for index < len(data) {
		split := strings.Fields(data[index])
		switch {
		case split[0] == "cpy":
			value, err := strconv.Atoi(split[1])
			dest = split[2]
			if err != nil {
				source = split[1]
				reg[dest] = reg[source]
			} else {
				reg[dest] = value
			}
		case split[0] == "inc":
			reg[split[1]] += 1
		case split[0] == "dec":
			reg[split[1]] -= 1
		case split[0] == "jnz":
			value, err := strconv.Atoi(split[1])
			jmp, _ := strconv.Atoi(split[2])
			if err != nil {
				if reg[split[1]] != 0 {
					index += jmp - 1
				}
			} else {
				if value != 0 {
					index += jmp - 1
				}
			}
		}
		index++
	}
	return reg["a"]
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 12 GO !\n")

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")

	fmt.Printf("Q1: The value inside reg a is : %d\n", day12Algo(data, false))
	fmt.Printf("Q1: The value inside reg a is : %d\n", day12Algo(data, true))

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
