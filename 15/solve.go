package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

type disk struct {
	startAt int
	pos     int
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 15 GO!\n")

	var disklist []disk
	var t int
	var part = 1

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")

	for _, e := range data {
		temp := strings.Replace(e, ".", "", -1)
		split := strings.Fields(temp)
		startAt, _ := strconv.Atoi(split[11])
		pos, _ := strconv.Atoi(split[3])
		disklist = append(disklist, disk{startAt, pos})
	}

wait:
	for i, disk := range disklist {
		if (disk.startAt+t+i+1)%disk.pos != 0 {
			t++
			goto wait
		}
	}
	fmt.Printf("Q%d: You need to press at time = %d\n", part, t)
	if part == 1 {
		part += 1
		disklist = append(disklist, disk{0, 11})
		goto wait
	}

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
