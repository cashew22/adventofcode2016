package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

var input = "abc"

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 14 GO (go routine experiment) !\n")

	index := make(chan int)
	baseIndex := make(chan int)
	hash := make(chan string)
	five := make(chan byte)

	//hash generator
	go func() {
		var hashMap map[int]string // lol
		hashMap = make(map[int]string)
		for {
			//fmt.println("hashGen wait for work")
			i := <-index
			//fmt.println("hashGen need to hash ", input, i)
			if hashMap[i] != "" {
				//fmt.println("Hash is in the list", hashMap[i])
				hash <- hashMap[i]
			} else {
				hashMap[i] = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", input, i))))
				//fmt.println("Hew hash", hashMap[i])
				hash <- hashMap[i]
			}
		}
	}()

	//Find 3 of a kind
	go func() {
		i := 0
		for {
			//fmt.println("3ofaKind send work to hashgen")
			index <- i
			h := <-hash
			//fmt.println("Receive hash from hashgen", h)
			for j := 0; j < len(h)-2; j++ {
				w := h[j : j+3]
				if w[0] == w[1] && w[1] == w[2] {
					//fmt.println("find 3ofaKind", string(w))
					baseIndex <- i
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
		//fmt.println("5ofaKind wait for work")
		i := <-baseIndex
		c := <-five
		//fmt.println("5ofaKind receive work", i, c)
		k := 0
		for k < 1000 {
			find := false
			index <- i + k
			h := <-hash
			//fmt.println("5ofaKind receive hash", i, h)
			for j := 0; j < len(h)-4; j++ {
				w := h[j : j+5]
				if w[0] == w[1] && w[1] == w[2] && w[2] == w[3] && w[3] == w[4] && w[4] == c {
					//fmt.println("find 5ofaKind", string(w))
					key = append(key, i)
					fmt.Println(len(key), i)
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
			fmt.Printf("Q1 :%d\n", key[63])
			break
		}
		i++
	}

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
