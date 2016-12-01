package main

import (
"fmt"
"os"
"bufio"
"time"
"strings"
"strconv"
"math"
)

type oper struct {
	direction string
	number int
}

type coordinate struct {
	x int
	y int
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func create_ops(input string) []oper {
	data := strings.Split(input, ",")
	var opers []oper
	var oper oper
	for _,e := range data {
		e = strings.TrimSpace(e)
		oper.direction = e[:1]
		num, err := strconv.Atoi(e[1:])
		check(err)
		oper.number = num
		opers = append(opers, oper)
	}

	return opers
}

func duplicate(list []coordinate) (coor coordinate, match bool) {
	encountered := map[coordinate]bool{}
	for e := range list {
		if encountered[list[e]] == true {
			return list[e], true
		} else {
			encountered[list[e]] = true
		}
	}

	return coordinate{0, 0}, false
}

func main() {
    start := time.Now() //timestamp
    fmt.Printf("Solution for day 18 GO !\n")

    var ops []oper
    var x = 0
    var y = 0
    var dir = "N"
    var day2_solve = false
    var day2_coor coordinate
    var list []coordinate
    temp := coordinate{0,0}
    match := false

    //Read file
    data, err := readLines("input.txt")
    check(err)

    ops = create_ops(data[0])

    list = append(list, coordinate{0,0})

    for _, e := range ops {
    	old_x := x
    	old_y := y

    	switch {
    	case dir == "N":
    		if e.direction == "L" {
    			x = x - e.number
    			dir = "E"
    			for i:= old_x; i > x; i-- { list = append(list, coordinate{i-1,y}) }
    			temp, match = duplicate(list)
    		} else {
    			x = x + e.number
    			dir = "W"
    			for i:= old_x; i < x; i++ { list = append(list, coordinate{i+1,y}) }
    			temp, match = duplicate(list)
    		}
    	case dir == "S":
    		if e.direction == "L" {
    			x = x + e.number
    			dir = "W"
    			for i:= old_x; i < x; i++ { list = append(list, coordinate{i+1,y}) }
    			temp, match = duplicate(list)
    		} else {
    			x = x - e.number
    			dir = "E"
    			for i:= old_x; i > x; i-- { list = append(list, coordinate{i-1,y}) }
    			temp, match = duplicate(list)
    		}
    	case dir == "W":
    		if e.direction == "L" {
    			y = y + e.number
    			dir = "N"
    			for i:= old_y; i < y; i++ { list = append(list, coordinate{x,i+1}) }
    			temp, match = duplicate(list)
    		} else {
    			y = y - e.number
    			dir = "S"
    			for i:= old_y; i > y; i-- { list = append(list, coordinate{x,i-1}) }
    			temp, match = duplicate(list)
    		}
    	case dir == "E":
    		if e.direction == "L" {
    			y = y - e.number
    			dir = "S"
    			for i:= old_y; i > y; i-- { list = append(list, coordinate{x,i-1}) }
    			temp, match = duplicate(list)
    		} else {
    			y = y + e.number
    			dir = "N"
    			for i:= old_y; i < y; i++ { list = append(list, coordinate{x,i+1}) }
    			temp, match = duplicate(list)
    		}
    	}

    	if !day2_solve && match {
    		day2_solve = true
    		day2_coor = temp
    	}

    }

    distance := math.Abs(0 - float64(x)) + math.Abs(0 - float64(y))
    distance_day2 := math.Abs(0 - float64(day2_coor.x)) + math.Abs(0 - float64(day2_coor.y))

    fmt.Printf("Q1 answer is : %d\n", int(distance))
    fmt.Printf("Q2 answer is : %d\n", int(distance_day2))
    //Elapse time
    elapsed := time.Since(start)
    fmt.Printf("Execution took %s\n", elapsed)
}
