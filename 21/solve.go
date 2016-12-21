package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func revers(lst []string) chan string {
    ret := make(chan string)
    go func() {
        for i, _ := range lst {
            ret <- lst[len(lst)-1-i]
        }
        close(ret)
    }()
    return ret
}


func swapPos(password, x_s, y_s string) string {
	x, _ := strconv.Atoi(x_s)
	y, _ := strconv.Atoi(y_s)
	pass := []rune(password[:])
	newPass := make([]rune, len(pass))
	copy(newPass, pass)
	newPass[x] = pass[y]
	newPass[y] = pass[x]
	return string(newPass)
}

func swapLet(password, x, y string) string {
	pass := []rune(password[:])
	var newPass string
	for i:= 0; i < len(password); i++ {
		if string(pass[i]) == x {
			newPass += y
		} else if string(pass[i]) == y {
			newPass += x
		} else {
			newPass += string(pass[i])
		}
	}
	return newPass
}

func reverse(password, x_s, y_s string) string {
	x, _ := strconv.Atoi(x_s)
	y, _ := strconv.Atoi(y_s)
	pass := []rune(password[:])
	newPass := ""
	rev := 0
	for i, e := range pass {
		if i >= x && i <= y {
			newPass += string(pass[y-rev]) 
			rev++
		} else {
			newPass += string(e)
		}
	}
	return newPass
}

func rotateLeft(password, step_s string) string {
	step, _ := strconv.Atoi(step_s)
	pass := []rune(password[:])
	passDouble := make([]rune, len(pass))
	newPass := ""
	copy(passDouble, pass)

	for _, e := range pass {
		passDouble = append(passDouble, e)
	}

	for i:=step; i < len(password) + step; i++ {
		newPass += string(passDouble[i])
	}
	return newPass
}

func rotateRight(password, step_s string) string {
	step, _ := strconv.Atoi(step_s)
	pass := []rune(password[:])
	passDouble := make([]rune, len(pass))
	newPass := ""
	copy(passDouble, pass)

	for _, e := range pass {
		passDouble = append(passDouble, e)
	}

	for i:=len(password) - step; i < len(password)*2 - step; i++ {
		newPass += string(passDouble[i])
	}
	return newPass
}

func rotateBased(password, step_s string, day2 bool) string {
	step := strings.Index(password, step_s)
	if step >= 4 {
		step++
	}
	step++
	step_s = strconv.Itoa(step)
	
	for i:= 0; i < step; i++ {
		if day2 {
			password = rotateLeft(password, "1")
		} else {
			password = rotateRight(password, "1")
		}
	}
	return password
}

func move(password, x_s, y_s string) string {
	x, _ := strconv.Atoi(x_s)
	y, _ := strconv.Atoi(y_s)
	pass := []rune(password[:])
	letter := pass[x]
	newPass:= ""
	pass = append(pass[:x], pass[x+1:]...)

	for i:= 0; i < y; i++ {
		newPass += string(pass[i])
 	}
 	newPass += string(letter)
 	for i:= y; i < len(password)-1; i++ {
		newPass += string(pass[i])
 	}
 	return string(newPass)
}

func scramble(password string, data []string) string {
	for _, ops := range data {
		fmt.Println(ops)
		split := strings.Fields(ops)
		if split[0] == "swap" && split[1] == "position" {
			password = swapPos(password, split[2], split[5])
		} else if split[0] == "swap" && split[1] == "letter" {
			password = swapLet(password, split[2], split[5])
		} else if split[0] == "reverse" {
			password = reverse(password, split[2], split[4])
		} else if split[0] == "rotate" && split[1] == "left" {
			password = rotateLeft(password, split[2])
		} else if split[0] == "rotate" && split[1] == "right" {
			password = rotateRight(password, split[2])
		} else if split[0] == "rotate" && split[1] == "based" {
			password = rotateBased(password, split[6], false)
		} else if split[0] == "move" {
			password = move(password, split[2], split[5])
		}
		fmt.Println(password)
	}
	return password
}

func unScramble(password string, data []string) string {
	for ops := range revers(data) {
		fmt.Println(ops)
		split := strings.Fields(ops)
		if split[0] == "swap" && split[1] == "position" {
			password = swapPos(password, split[5], split[2])
		} else if split[0] == "swap" && split[1] == "letter" {
			password = swapLet(password, split[5], split[2])
		} else if split[0] == "reverse" {
			password = reverse(password, split[2], split[4])
		} else if split[0] == "rotate" && split[1] == "left" {
			password = rotateRight(password, split[2])
		} else if split[0] == "rotate" && split[1] == "right" {
			password = rotateLeft(password, split[2])
		} else if split[0] == "rotate" && split[1] == "based" {
			password = rotateBased(password, split[6], true)
		} else if split[0] == "move" {
			password = move(password, split[5], split[2])
		}
		fmt.Println(password)
	}
	return password
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 21 GO!\n")

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")

	fmt.Printf("Q1: The scramble password is: %s\n", scramble("abcde", data))
	fmt.Printf("Q2: The unscramble password is: %s\n", unScramble("decab", data))

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
