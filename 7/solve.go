package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

func supportTLS(ip string) bool {
	var (
		support_TLS bool
		valid       bool = true
	)
	for i := 0; i+3 < len(ip); i++ { //The goal is to stop at the last window
		window := []rune(ip[i : i+4])

		if strings.Contains(string(window), "[") { // interring invalid zone
			valid = false
		} else if strings.Contains(string(window), "]") { // exiting invalid zone
			valid = true
		}

		match := window[0] == window[3] && window[1] == window[2] && window[0] != window[1]
		if match && !valid {
			return false
		} else if match && valid {
			support_TLS = true
		}
	}
	return support_TLS

}

func supportSSL(ip string) bool {
	var aba_zone = true

	for i := 0; i+2 < len(ip); i++ { //The goal is to stop at the last window
		window := []rune(ip[i : i+3])

		if strings.Contains(string(window), "[") { // interring BAB zone
			aba_zone = false
		} else if strings.Contains(string(window), "]") { // exiting BAB zone
			aba_zone = true
		}

		match := window[0] == window[2] && aba_zone
		if match { //we found a ABA, search for BAB inside []
			bab_zone := false
			for i := 0; i < len(ip)-3; i++ {
				bab_window := []rune(ip[i : i+3])
				if strings.Contains(string(bab_window), "[") { // interring BAB zone
					bab_zone = true
				} else if strings.Contains(string(bab_window), "]") { // exiting invalid zone
					bab_zone = false
				}

				match_bab := bab_window[0] == window[1] && bab_window[1] == window[0] &&
					bab_window[1] == window[2] && bab_window[2] == window[1] && bab_zone
				if match_bab {
					return true
				}
			}
		}
	}
	return false
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 7 GO !\n")

	var (
		support_TLS int
		support_SSL int
	)

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")

	for _, ip := range data {
		if supportTLS(ip) {
			support_TLS++
		}
		if supportSSL(ip) {
			support_SSL++
		}
	}

	fmt.Printf("Q1: Number of IP supporting TLS = %d\n", support_TLS)
	fmt.Printf("Q2: Number of IP supporting SSL = %d\n", support_SSL)

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
