package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

type Bot struct {
	number   int
	nbOfChip int
	low      int
	high     int
}

type Move struct {
	number   int
	low      int
	high     int
	lowdest  string
	highDest string
}

func parseValue(input string) (int, int) {
	var chip, number int
	fmt.Sscanf(input, "value %d goes to bot %d", &chip, &number)
	return chip, number
}

func parseMove(input string) Move {
	var number, low, high int
	var lowdest, highDest string
	fmt.Sscanf(input, "bot %d gives low to %s %d and high to %s %d", &number, &lowdest,
		&low, &highDest, &high)
	return Move{number, low, high, lowdest, highDest}
}

func getBot(botList []Bot, number int) (Bot, error) {
	var bot Bot
	for _, bot := range botList {
		if bot.number == number {
			return bot, nil
		}
	}
	return bot, errors.New("No Bot")
}

func updateBotList(botList []Bot, bot Bot) []Bot {
	for i := 0; i < len(botList); i++ {
		if botList[i].number == bot.number {
			botList[i] = bot
			return botList
		}
	}
	botList = append(botList, bot)
	return botList
}

func getBotReady(botList []Bot) (Bot, error) {
	var bot Bot
	for _, bot := range botList {
		if bot.nbOfChip == 2 {
			return bot, nil
		}
	}
	return bot, errors.New("No Bot")
}

func assignChip(bot *Bot, chip int) {
	if chip > bot.high {
		bot.low = bot.high
		bot.high = chip
	} else {
		bot.low = chip
	}
	bot.nbOfChip++
}

func findMove(data []string, bot Bot) (Move, error) {
	var move Move
	for _, input := range data {
		if strings.HasPrefix(input, "bot") {
			move = parseMove(input)
			if move.number == bot.number {
				return move, nil
			}
		}
	}
	return move, errors.New("No move")
}

func main() {
	start := time.Now() //timestamp
	fmt.Printf("Solution for day 10 GO !\n")

	var (
		botList   []Bot
		output1   []int
		output2   []int
		output3   []int
		botToKeep = 999999
	)

	//Read file
	file, _ := ioutil.ReadFile("input.txt")
	data := strings.Split(string(file), "\n")

	for _, e := range data {
		if strings.HasPrefix(e, "value") {
			chip, number := parseValue(e)
			bot, err := getBot(botList, number)
			if err != nil {
				bot = Bot{number, 0, 0, 0}
			}
			assignChip(&bot, chip)
			botList = updateBotList(botList, bot)
		}
	}

	for {
		bot, err := getBotReady(botList)
		if err != nil {
			fmt.Println("No more bot ready to run. Done")
			break
		}
		move, _ := findMove(data, bot)
		if bot.high == 61 && bot.low == 17 {
			botToKeep = bot.number
		}

		if move.lowdest == "output" {
			if move.low == 0 {
				output1 = append(output1, bot.low)
			} else if move.low == 1 {
				output2 = append(output2, bot.low)
			} else if move.low == 2 {
				output3 = append(output3, bot.low)
			}
		} else {
			botLow, err := getBot(botList, move.low)
			if err != nil {
				botLow = Bot{move.low, 0, 0, 0}
			}
			assignChip(&botLow, bot.low)
			botList = updateBotList(botList, botLow)
		}

		if move.highDest == "output" {
			if move.high == 0 {
				output1 = append(output1, bot.high)
			} else if move.high == 1 {
				output2 = append(output2, bot.high)
			} else if move.high == 2 {
				output3 = append(output3, bot.high)
			}
		} else {
			botHigh, err := getBot(botList, move.high)
			if err != nil {
				botHigh = Bot{move.high, 0, 0, 0}
			}
			assignChip(&botHigh, bot.high)
			botList = updateBotList(botList, botHigh)
		}
		bot.high = 0
		bot.low = 0
		bot.nbOfChip = 0
		botList = updateBotList(botList, bot)
	}

	fmt.Println(output1, output2, output3)
	fmt.Printf("Q1: Bot number %d is responsible of comparing 17 and 61\n",
		botToKeep)
	fmt.Printf("Q2: The answer is %d\n", output1[0]*output2[0]*output3[0])

	//Elapse time
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s\n", elapsed)
}
