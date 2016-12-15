package main

import (
	"crypto/md5"
	"fmt"
	"sync"
	"time"
)

var input = "zpqevtbw"
var hashMap map[int]string // lol
var mutex = &sync.Mutex{}

func getHash(i int) string {
	mutex.Lock()
	if hashMap[i] == "" {
		hashMap[i] = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
	}
	mutex.Unlock()
	return hashMap[i]
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 14 GO part 1(go routine experiment) !\n")

	index := make(chan int)
	five := make(chan byte)
	hashMap = make(map[int]string)

	//Find 3 of a kind
	go func() {
		i := 0
		for {
			h := getHash(i)
			for j := 0; j < len(h)-2; j++ {
				w := h[j : j+3]
				if w[0] == w[1] && w[1] == w[2] {
					index <- i
					five <- w[0]
					break
				}
			}
			i++
		}
	}()

	//find 5 of a kind in next 1000
	var key []int
	for {
		i := <-index
		c := <-five
		k := 1
		for k < 1001 {
			find := false
			h := getHash(i + k)
			for j := 0; j < len(h)-4; j++ {
				w := h[j : j+5]
				if w[0] == w[1] && w[1] == w[2] && w[2] == w[3] && w[3] == w[4] && w[4] == c {
					key = append(key, i)
					fmt.Println(len(key), i, i+k)
					find = true
					break
				}
			}
			if find {
				break
			}
			k++
		}
		if len(key) == 64 {
			fmt.Printf("Q1: The 64th key is %d\n", key[63])
			break
		}
		i++
	}
	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
