package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var (
	wide = 50
	tall = 6
)

func rect(screen [6][50]string, params string) [6][50]string {
	splitted := strings.Split(params, "x")
	y, _ := strconv.Atoi(splitted[0])
	x, _ := strconv.Atoi(splitted[1])

	for i := 0; i < y; i++ {
		for k := 0; k < x; k++ {
			screen[k][i] = "#"
		}
	}
	return screen
}

func rotate(screen [6][50]string, ops string, index_s string, arg string) [6][50]string {
	splitted := strings.Split(index_s, "=")
	index, _ := strconv.Atoi(splitted[1])
	shift, _ := strconv.Atoi(arg)

	if ops == "row" {
		row_copy := screen[index]
		for i := 0; i < wide; i++ {
			if i-shift < 0 {
				screen[index][i] = row_copy[wide-(shift-i)]
			} else {
				screen[index][i] = row_copy[i-shift]
			}
		}
	} else {
		var column [6]string
		for i, row := range screen {
			column[i] = row[index]
		}

		column_copy := column
		for i := 0; i < tall; i++ {
			if i-shift < 0 {
				column[i] = column_copy[tall-(shift-i)]
			} else {
				column[i] = column_copy[i-shift]
			}
		}

		for i := 0; i < tall; i++ {
			screen[i][index] = column[i]
		}
	}
	return screen
}

func printLCD(screen [6][50]string) {
	for _, row := range screen {
		fmt.Println(row)
	}
}

func light_on(screen [6][50]string) int {
	var lightOn = 0
	for _, row := range screen {
		for _, light := range row {
			if light == "#" {
				lightOn++
			}
		}
	}
	return lightOn
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 8 GO !\n")

	var screen [6][50]string
	for i := 0; i < tall; i++ {
		for j := 0; j < wide; j++ {
			screen[i][j] = " "
		}
	}

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")

	for _, ops := range data {
		splitted := strings.Fields(ops)

		switch splitted[0] {
		case "rect":
			screen = rect(screen, splitted[1]) //rect 3x2
		case "rotate":
			screen = rotate(screen, splitted[1], splitted[2], splitted[4]) //rotate row y=0 by 4
		}
	}

	fmt.Printf("Q1: Number of light on = %d\n", light_on(screen))
	printLCD(screen)

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
