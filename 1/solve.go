package main

import (
	"fmt"
    "os"
    "bufio"
    "time"
    "strings"
    "strconv"
)

type oper struct {
    direction string
    number int
}

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

func create_ops(input string) []oper {
    data := strings.Split(input, ",")
    var opers []oper
    var oper oper
    for _,e := range data {
        e = strings.TrimSpace(e)
        oper.direction = e[:1]
        num, err := strconv.Atoi(e[1:])
        check(err)
        oper.number = num
        opers = append(opers, oper)
    }

    return opers
}

func main() {
    start := time.Now() //timestamp
	fmt.Printf("Solution for day 18 GO !\n")

    var ops []oper

    //Read file
    data, err := readLines("input.txt")
    check(err)

    ops = create_ops(data[0])

    for _, e := range ops

    fmt.Println(ops)

    //Elapse time
    elapsed := time.Since(start)
    fmt.Printf("Execution took %s\n", elapsed)
}
