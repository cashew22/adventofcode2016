package main

import (
	"fmt"
    "io/ioutil"
    "time"
    "strings"
)

func valid(a int, b int, c int) bool {
    return a + b > c && a + c > b && c + b > a
}


func main() {
    start := time.Now() //timestamp
	fmt.Printf("Solution for day 18 GO !\n")

    //Read file
    file, _ := ioutil.ReadFile("input.txt")
    data := strings.Split(string(file), "\n")

    var valid_triangle = 0
    var valid_triangle_day2 = 0
    var a, b, c int

    for _, e := range data {
        fmt.Sscanf(e, "%d%d%d", &a, &b, &c)
        if valid(a, b, c) {valid_triangle++}
    }
    fmt.Printf("Q1: number of valid triangle = %d\n", valid_triangle)
    
    var a1, b1, c1, a2, b2, c2, a3 ,b3, c3 int 
    for i:=0; i<len(data); i+=3 {
        fmt.Sscanf(data[i], "%d%d%d", &a1, &a2, &a3)
        fmt.Sscanf(data[i+1], "%d%d%d", &b1, &b2, &b3)
        fmt.Sscanf(data[i+2], "%d%d%d", &c1, &c2, &c3)
        if valid(a1, b1, c1) {valid_triangle_day2++}
        if valid(a2, b2, c2) {valid_triangle_day2++}
        if valid(a3, b3, c3) {valid_triangle_day2++}

    }
    fmt.Printf("Q2: number of valid triangle = %d\n", valid_triangle_day2)

    //Elapse time
    elapsed := time.Since(start)
    fmt.Printf("Execution took %s\n", elapsed)
}
