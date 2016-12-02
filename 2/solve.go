package main

import (
	"fmt"
    "os"
    "bufio"
    "time"
)

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

func out_of_bound(pos int, ops string, day2 bool) bool {
	if !day2 {
		if ops == "D" {
			if pos == 7 || pos == 8 || pos == 9 {return true}
		} else if (ops == "U") {
			if pos == 1 || pos == 2 || pos == 3 {return true}
		} else if (ops == "L") {
			if pos == 7 || pos == 4 || pos == 1 {return true}
		} else {
			if pos == 9 || pos == 6 || pos == 3 {return true}
		}
	} else {
		if ops == "D" {
			if pos == 5 || pos == 9 || pos == 10 || pos == 12 || pos == 13 {return true}
		} else if (ops == "U") {
			if pos == 5 || pos == 9 || pos == 2 || pos == 4 || pos == 1 {return true}
		} else if (ops == "L") {
			if pos == 1 || pos == 2 || pos == 5 || pos == 10 || pos == 13 {return true}
		} else {
			if pos == 1 || pos == 4 || pos == 9 || pos == 12 || pos == 13 {return true}
		}
	}

	return false
}

func main() {
    start := time.Now() //timestamp
	fmt.Printf("Solution for day 18 GO !\n")

	var key_list []int
	pos := 5

    //Read file
    data, err := readLines("input.txt")
    check(err)

    //Q1
    for _,seq := range data {
    	for _, move := range seq {
    		if !out_of_bound(pos, string(move), false) {
    			if string(move) == "D" {
					pos = pos + 3
				} else if string(move) == "U" {
					pos = pos - 3
				} else if string(move) == "L" {
					pos = pos - 1
				} else {
					pos = pos + 1
				}
			}
		}
		key_list = append(key_list, pos)
    }

    fmt.Printf("Q1 answer is: ")
    for _,e := range key_list {
    	fmt.Printf("%d",e)
    }
    fmt.Printf("\n")

    var key_list_2 []int
    var row = 3
    //Q2
    for _,seq := range data {
    	for _, move := range seq {
    		if !out_of_bound(pos, string(move), true) {
    			if string(move) == "D" {
					if row == 1 || row == 4 {
						pos = pos + 2
					} else {
						pos = pos + 4
					}
					row++
				} else if string(move) == "U" {
					if row == 5 || row == 2 {
						pos = pos - 2
					} else {
						pos = pos - 4
					}
					row--
				} else if string(move) == "L" {
					pos = pos - 1
				} else {
					pos = pos + 1
				}
			}
		}
		key_list_2 = append(key_list_2, pos)
    }

    fmt.Printf("Q2 answer is: ")
    for _,e := range key_list_2 {
    	if e > 9 {
    		if e == 10 {fmt.Printf("A")}
    		if e == 11 {fmt.Printf("B")}
    		if e == 12 {fmt.Printf("C")}
    		if e == 13 {fmt.Printf("D")}
    	} else {
    		fmt.Printf("%d",e)
    	}
    }
    fmt.Printf("\n")

    //Elapse time
    elapsed := time.Since(start)
    fmt.Printf("Execution took %s\n", elapsed)
}
