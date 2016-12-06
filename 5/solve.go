package main

import (
	"fmt"
    "time"
    "crypto/md5"
    "strings"
    "strconv"
)

func solveDay5(input string, day2 bool) string {
	var charFound = 0
	var password string
	var password_2 [8]string

	for i:=0; charFound < 8;i++ {
		input := fmt.Sprintf("%s%d", input, i)
		h := md5.Sum([]byte(input))
		hash := fmt.Sprintf("%x", h)

		if strings.HasPrefix(hash, "00000") {
			if !day2 { //Day1
				var sixth string
				for k,c := range hash {
					if k == 5 {
						sixth = string(c)
						break
					}
				}
				password += sixth
				charFound++
			} else { //Day2
				var pos int
				var key string
				var err error
				for k, c := range hash {
					if k == 5 { pos, err = strconv.Atoi(string(c)) }
					if k == 6 {
						key = string(c)
						break;
					}
				}
				if err == nil && pos >= 0 && pos < 8 && password_2[pos] == "" {
					password_2[pos] = key
					charFound++
				}
			} 
		}
	}

	if day2 {
		for _, e := range password_2 {
			password += e
		}
	}
	return password
}

func main() {
    start := time.Now() //timestamp
	fmt.Printf("Solution for day 5 GO !\n")

	var input = "uqwqemis"
	done := make(chan bool)

	go func() {
		pass_1 := solveDay5(input, false)
		fmt.Printf("Q1: %s\n", pass_1)
		done <- true
	}()
	go func() {
		pass_2 := solveDay5(input, true)
		fmt.Printf("Q2: %s\n", pass_2)
		done <- true
	}()

	for i := 0; i < 2; i++ {
        <-done
    }

    //Elapse time
    elapsed := time.Since(start)
    fmt.Printf("Execution took %s\n", elapsed)
}
