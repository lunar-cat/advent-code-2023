package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	scratchCards()
}

func scratchCards() {
	file, err := os.Open("./day-4/input.txt")
	if err != nil {
		log.Fatal("Missing File")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	winnerPoints := 0
	cards := 0
	i := 0
	repeatMap := make(map[int]int)
	for scanner.Scan() {
		i++
		repeatMap[i]++
		p, c := getWinnerPoints(scanner.Text())
		winnerPoints += p
		countRepetitions(repeatMap, c, i)
	}

	for _, v := range repeatMap {
		cards += v
	}
	fmt.Println(winnerPoints, cards)
}

func getWinnerPoints(line string) (int, int) {
	points := 0
	cards := 0
	numbers := strings.Split(line, "|")
	r := regexp.MustCompile(`\d+`)
	winningNumbers := r.FindAllString(numbers[0], -1)[1:]
	cardNumbers := r.FindAllString(numbers[1], -1)
	winningNumbersMap := make(map[string]bool)

	for _, numStr := range winningNumbers {
		winningNumbersMap[numStr] = true
	}

	for _, numStr := range cardNumbers {
		if winningNumbersMap[numStr] {
			cards++
			if points == 0 {
				points += 1
			} else {
				points *= 2
			}
		}
	}
	return points, cards
}

func countRepetitions(repeatMap map[int]int, cards int, currentCard int) {
	for j := 0; j < repeatMap[currentCard]; j++ {
		for m := 1; m <= cards; m++ {
			repeatMap[currentCard+m]++
		}
	}
}
