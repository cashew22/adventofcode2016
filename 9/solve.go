package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func processMarker(data string, day2 bool) int {
	var (
		markerZone = false
		fileLen    = 0
		numOfChar  = 0
		numOfIter  = 0
		marker     = ""
	)
	for i := 0; i < len(data); i++ {
		if string(data[i]) == "(" {
			markerZone = true
		} else if string(data[i]) == ")" {
			markerZone = false
			fmt.Sscanf(marker, "%dx%d", &numOfChar, &numOfIter)
			if day2 {
				fileLen += numOfIter * processMarker(data[i+1:i+numOfChar+1], day2)
			} else {
				fileLen += numOfChar * numOfIter
			}
			i += numOfChar
			marker = ""
		} else if markerZone {
			marker += string(data[i])
		} else {
			fileLen++
		}
	}

	return fileLen
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 8 GO !\n")

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")
	data[0] = strings.Replace(data[0], " ", "", -1)
	compressData := data[0][:]

	fmt.Printf("Q1: Len of decompressed data = %d\n", processMarker(compressData, false))
	fmt.Printf("Q2: Len of decompressed data = %d\n", processMarker(compressData, true))

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
