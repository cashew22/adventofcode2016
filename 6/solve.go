package main

import (
	"fmt"
    "io/ioutil"
    "time"
    "strings"
)

func main() {
    start := time.Now() //timestamp
	fmt.Printf("Solution for day 6 GO !\n")

    var m map[rune]int
    var passwordMax, passwordMin string

    //Read file
    file, _ := ioutil.ReadFile("input.txt")
    data := strings.Split(string(file), "\n")

    lenght := len(data[0])

    for i:=0; i < lenght; i++ {
        m = make(map[rune]int)
        for _, s := range data { //loop over lines
            for j, c := range s { //loop over line to find rune
                if j == i {
                 m[c]++
                 break
                }
            }
        }
        //Find the most and he least rune in eac position
        var keepMax, keepMin rune
        var lastKeyMax = 0
        var lastKeyMin = 10000
        for k, v := range m { 
            if v > lastKeyMax {
                keepMax = k
                lastKeyMax = v
            }
            if v < lastKeyMin {
                keepMin = k
                lastKeyMin = v
            }
        }
        passwordMax += string(keepMax)
        passwordMin += string(keepMin)
    }

    fmt.Printf("Q1: The password is : %s\n", passwordMax)
    fmt.Printf("Q2: The password is : %s\n", passwordMin)

    //Elapse time
    elapsed := time.Since(start)
    fmt.Printf("Execution took %s\n", elapsed)
}
