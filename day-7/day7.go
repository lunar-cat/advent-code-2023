package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const PartTwo = true

type Hand struct {
	cards    string
	bid      int
	strength int
}

func main() {
	camelCards()
}

func camelCards() {
	file, err := os.Open("./day-7/input.txt")
	if err != nil {
		log.Fatal("Missing File")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var hands []Hand
	for scanner.Scan() {
		hands = append(hands, processHand(scanner.Text()))
	}
	sortHands(hands)
	totalWinnings := getTotalWinnings(hands)
	fmt.Println(totalWinnings)
}

func processHand(line string) Hand {
	parts := strings.Split(line, " ")

	cards := parts[0]
	bid, _ := strconv.Atoi(parts[1])
	strength := calculateStrength(cards)
	return Hand{cards, bid, strength}
}

func calculateStrength(cards string) int {
	counter := make(map[string]int)
	var five []int
	var four []int
	var three []int
	var pair []int
	var jokers []int
	for _, char := range cards {
		if PartTwo && string(char) == "J" {
			jokers = append(jokers, 1)
			continue
		}
		counter[string(char)]++
	}
	for _, v := range counter {
		switch v {
		case 5:
			five = append(five, 5)
		case 4:
			four = append(four, 4)
		case 3:
			three = append(three, 3)
		case 2:
			pair = append(pair, 2)
		}
	}

	switch {
	case len(five) == 1 || len(jokers) == 5 || len(jokers) == 4 || (len(four) == 1 && len(jokers) == 1) || (len(three) == 1 && len(jokers) == 2) || (len(pair) == 1 && len(jokers) == 3):
		return 6
	case len(four) == 1 || len(jokers) == 3 || (len(three) == 1 && len(jokers) == 1) || (len(pair) == 1 && len(jokers) == 2):
		return 5
	case len(three) == 1 && len(pair) == 1 || (len(three) == 1 && len(jokers) == 1) || (len(pair) == 2 && len(jokers) == 1):
		return 4
	case len(three) == 1 || len(jokers) == 2 || (len(pair) == 1 && len(jokers) == 1):
		return 3
	case len(pair) == 2 || (len(pair) == 1 && len(jokers) == 1):
		return 2
	case len(pair) == 1 || len(jokers) == 1:
		return 1
	default:
		return 0
	}
}

func sortHands(hands []Hand) {
	strengthMap := map[string]int{
		"A": 13,
		"K": 12,
		"Q": 11,
		"J": 10,
		"T": 9,
		"9": 8,
		"8": 7,
		"7": 6,
		"6": 5,
		"5": 4,
		"4": 3,
		"3": 2,
		"2": 1,
	}
	if PartTwo {
		strengthMap["J"] = 0
	}
	sort.SliceStable(hands, func(i, j int) bool {
		if hands[i].strength == hands[j].strength {
			for k := 0; k < len(hands[j].cards); k++ {
				leftCard := strengthMap[string(hands[i].cards[k])]
				rightCard := strengthMap[string(hands[j].cards[k])]
				if leftCard == rightCard {
					continue
				}
				return leftCard < rightCard
			}
		}
		return hands[i].strength < hands[j].strength
	})
}

func getTotalWinnings(hands []Hand) int {
	total := 0
	for i := 1; i <= len(hands); i++ {
		total += hands[i-1].bid * i
	}
	return total
}
