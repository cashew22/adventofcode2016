package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"strconv"
)

var (
	wide = 6
	tall = 3
	)

func rect(screen [3][6]bool, params string) [3][6]bool {
	splitted := strings.Split(params, "x")
	y, _ := strconv.Atoi(splitted[0])
	x, _ := strconv.Atoi(splitted[1])
	fmt.Println(y, x)

	for i:=0; i<y; i++ {
		for k:=0; k<x; k++ {
			screen[k][i] = true
		}
	}
	return screen
}

func rotate(screen [3][6]bool, ops string, index_s string, arg string) [3][6]bool {
	splitted := strings.Split(index_s, "=")
	index, _ := strconv.Atoi(splitted[1])
	shift, _ := strconv.Atoi(arg)
	if ops == "row" {
		row_copy := bool[index]
		for i:=0; i < wide; i++ {
			if (i-shift < 0){
				screenp[index][i] = row_copy[]
			}
			screen[index][i] = row_copy[i-shift]
		}
	} else {

	}
	return screen
}

func printLCD(screen [3][6]bool) {
	fmt.Println("---------------------")
	for _, row := range screen {
		fmt.Println(row)
	}
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 8 GO !\n")

	var screen [3][6]bool

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")

	for _, ops := range data {
		splitted := strings.Fields(ops)

		switch splitted[0] {
		case "rect" :
			screen = rect(screen, splitted[1]) //rect 3x2
		case "rotate":
			screen = rotate(screen, splitted[1], splitted[2], splitted[4]) //rotate row y=0 by 4
		}
		printLCD(screen)
	}

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}