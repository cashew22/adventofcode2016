package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

var (
	end_x    = 3
	end_y    = 3
	input    = "vwbaicqe"
	offset   = len(input)
	pathList []string
	LenMin   = 1000000000000
	LenMax   = 0
)

func getDoor(input string) string {
	hash := md5.Sum([]byte(fmt.Sprintf("%s", input)))
	return hex.EncodeToString(hash[:4])
}

func isOpen(doors string, dir string) bool {
	if dir == "U" {
		return doors[0] == 'b' || doors[0] == 'c' || doors[0] == 'd' || doors[0] == 'e' ||
			doors[0] == 'f'
	}
	if dir == "D" {
		return doors[1] == 'b' || doors[1] == 'c' || doors[1] == 'd' || doors[1] == 'e' ||
			doors[1] == 'f'
	}
	if dir == "L" {
		return doors[2] == 'b' || doors[2] == 'c' || doors[2] == 'd' || doors[2] == 'e' ||
			doors[2] == 'f'
	}
	if dir == "R" {
		return doors[3] == 'b' || doors[3] == 'c' || doors[3] == 'd' || doors[3] == 'e' ||
			doors[3] == 'f'
	}
	return false
}

func isWall(x, y int, dir string) bool {
	if dir == "D" {
		return y == 3
	}
	if dir == "R" {
		return x == 3
	}
	if dir == "U" {
		return y == 0
	}
	if dir == "L" {
		return x == 0
	}
	return false
}

func process(x, y int, input string, last string) {
	if x == end_x && y == end_y {
		pathList = append(pathList, input[offset:])
		if LenMin > len(input) {
			LenMin = len(input)
		}
		if LenMax < len(input) {
			LenMax = len(input)
		}
		return
	}

	doors := getDoor(input)
	if isOpen(doors, "D") && !isWall(x, y, "D") && last != "U" {
		process(x, y+1, input+"D", "D")
	}
	if isOpen(doors, "R") && !isWall(x, y, "R") && last != "L" {
		process(x+1, y, input+"R", "R")
	}
	if isOpen(doors, "U") && !isWall(x, y, "U") && last != "D" {
		process(x, y-1, input+"U", "U")
	}
	if isOpen(doors, "L") && !isWall(x, y, "L") && last != "R" {
		process(x-1, y, input+"L", "L")
	}

	//try to go back
	if last == "U" && isOpen(doors, "D") {
		process(x, y+1, input+"D", "D")
	}
	if last == "D" && isOpen(doors, "U") {
		process(x, y-1, input+"U", "U")
	}
	if last == "L" && isOpen(doors, "R") {
		process(x+1, y, input+"R", "R")
	}
	if last == "R" && isOpen(doors, "L") {
		process(x-1, y, input+"L", "L")
	}
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 17 GO!\n")

	process(0, 0, input, "")
	for _, e := range pathList {
		if len(e)+offset == LenMin {
			fmt.Printf("Q1: Smallest path is %s\n", e)
		}
	}
	for _, e := range pathList {
		if len(e)+offset == LenMax {
			fmt.Printf("Q2: Longest path is %d steps long\n", len(e))
			break
		}
	}

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
