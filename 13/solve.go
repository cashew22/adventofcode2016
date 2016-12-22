package main

import (
	"fmt"
	"time"
)

type coor struct {
	y int
	x int
}

var (
	maze   [50][50]string
	borne  = 50
	input  = 1350
	target = coor{39, 31}
)

func popcount(x uint64) int {
	// bit population count, see
	// http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
	x -= (x >> 1) & 0x5555555555555555
	x = (x>>2)&0x3333333333333333 + x&0x3333333333333333
	x += x >> 4
	x &= 0x0f0f0f0f0f0f0f0f
	x *= 0x0101010101010101
	return int(x >> 56)
}

func fill(y, x int) string {
	if y == 0 || y == borne-1 || x == 0 || x == borne-1 {
		return "#"
	}
	if popcount(uint64(x*x+3*x+2*x*y+y+y*y+input))%2 == 0 {
		return "."
	} else {
		return "#"
	}
}

func printMaze() {
	for y := 0; y < borne; y++ {
		fmt.Printf("%d ", y)
		if y == 0 {
			for x := 0; x < borne; x++ {
				fmt.Printf("%d", x%10)
			}
			fmt.Printf("\n  ")
		}
		for x := 0; x < borne; x++ {
			if x == target.x && y == target.y {
				fmt.Printf("X")
			} else {
				fmt.Printf(maze[y][x])
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 13 GO !\n")

	for y := 0; y < borne; y++ {
		for x := 0; x < borne; x++ {
			maze[y][x] = fill(y, x)
		}
	}

	printMaze()

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
