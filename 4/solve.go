package main

import (
	"fmt"
    "io/ioutil"
    "time"
    "strings"
    "strconv"
    "sort"
)

type char struct {
	letter string
	occur int
}

type char_list []char

func (list char_list) Len() int {
	return len(list)
}

func (list char_list) Less(i, j int) bool {
	if list[i].occur > list[j].occur {return true}
	if list[i].occur < list[j].occur {return false}
	return list[i].letter < list[j].letter
}

func (list char_list) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func contains(input char, list []char) bool {
	for _, e := range list {
		if input.letter == e.letter {return true}
	}
	return false
}


func isReal(input string, checksum string) bool {
	var s char_list
	var calcChecksum = ""

	input = strings.Replace(input, "-", "", -1)
	
	for _, e := range input {
		count := strings.Count(input, string(e))
		c := char{string(e), count}
		if !contains(c, s) { s = append(s, c) }
	}

	sort.Sort(s)

	for i:=0; i < 5;i++ {
		calcChecksum += s[i].letter
	}

	return calcChecksum == checksum
}

func shift(r rune, shift int) rune {
	s := r + rune(shift % 26)
    if s > 'z' { return s - 26 }
    if s < 'a' { return s + 26 }
    return s
}

func decode(input string, fieldID int) string {
	var retval []rune
	for _, e := range input {
		if e == '-' {
			retval = append(retval, ' ')
		} else {
			retval = append(retval, shift(e, fieldID))
		}
	}
	return string(retval)
}

func main() {
    start := time.Now() //timestamp
	fmt.Printf("Solution for day 4 GO !\n")

	var sumFieldId = 0
	var npObject = 0

    file, _ := ioutil.ReadFile("input.txt")
    data := strings.Split(string(file), "\n")

    for _, e := range data {
        checksum := e[len(e)-6:len(e)-1]
        input := e[:len(e)-10]
        fieldID, _ := strconv.Atoi(e[len(e)-10:len(e)-7])
        if isReal(input, checksum) {
        	sumFieldId += fieldID
        	if strings.Contains(decode(string(input), fieldID), "northpole") {
        		npObject = fieldID
        	}
        }
    }
    fmt.Printf("Q1: The sum of real room field id is: %d\n", sumFieldId)
    fmt.Printf("Q2: North pole objects are in room: %d\n", npObject)

    //Elapse time
    elapsed := time.Since(start)
    fmt.Printf("Execution took %s\n", elapsed)
}
