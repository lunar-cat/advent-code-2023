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
	for scanner.Scan() {
		winnerPoints += getWinnerPoints(scanner.Text())
	}
	fmt.Println(winnerPoints)
}

func getWinnerPoints(line string) int {
	points := 0
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
			if points == 0 {
				points += 1
			} else {
				points *= 2
			}
		}
	}
	return points
}
