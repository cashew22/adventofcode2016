package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 8 GO !\n")

	var (
		state          = "normal"
		decompressData = ""
		marker         = ""
		buffer         = ""
		iter           int
		numOfChar      int
	)

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")
	data[0] = strings.Replace(data[0], " ", "", -1)

	for _, char := range data[0] {
		switch state {
		case "normal":
			buffer = ""
			numOfChar = 0
			iter = 0
			if char == '(' {
				state = "newMarker"
			} else {
				decompressData += string(char)
			}

		case "newMarker":
			if char == ')' {
				state = "DecompressStart"
			} else {
				marker += string(char)
			}

		case "DecompressStart":
			splitted := strings.Split(marker, "x")
			marker = ""
			numOfChar, _ = strconv.Atoi(splitted[0])
			iter, _ = strconv.Atoi(splitted[1])
			buffer += string(char)
			if numOfChar == 1 {
				for i := 0; i < iter; i++ {
					decompressData += buffer
				}
				if char == '(' {
					state = "newMarker"
				} else {
					state = "normal"
				}
			} else {
				state = "Decompress"
			}

		case "Decompress":
			if len(buffer) == numOfChar-1 {
				buffer += string(char)
				for i := 0; i < iter; i++ {
					decompressData += buffer
				}
				if char == '(' {
					state = "newMarker"
				} else {
					state = "normal"
				}
			} else {
				buffer += string(char)
			}
		}
	}
	fmt.Printf("Q1: Len of decompressed data = %d\n", len(decompressData))

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
