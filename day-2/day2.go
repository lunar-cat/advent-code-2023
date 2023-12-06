package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	games()
}

func games() {
	file, err := os.Open("./day-2/input.txt")
	if err != nil {
		log.Fatal("Missing File")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sumID, sumPower := getPossibleGames(scanner)
	fmt.Println(sumID, sumPower)
}

func getPossibleGames(scanner *bufio.Scanner) (int, int) {
	gameConstraint := make(map[string]int)
	gameConstraint["red"] = 12
	gameConstraint["green"] = 13
	gameConstraint["blue"] = 14
	sumId := 0
	sumPower := 0
	for scanner.Scan() {
		id, power := getGameData(scanner.Text(), gameConstraint)
		sumId += id
		sumPower += power
	}
	return sumId, sumPower
}

func getGameData(line string, gameConstraint map[string]int) (int, int) {
	possible := true
	gameCount := make(map[string]int)

	gameData := strings.Split(line, ":")
	gameHeader := strings.Split(strings.TrimSpace(gameData[0]), " ")
	gameID, err := strconv.Atoi(gameHeader[1])
	if err != nil {
		log.Fatalf("Found non numeric ID: %s", gameHeader[1])
	}

	for _, groups := range strings.Split(gameData[1], ";") {
		groupData := make(map[string]int)
		for _, group := range strings.Split(groups, ",") {
			game := strings.Split(strings.TrimSpace(group), " ")
			amount, err := strconv.Atoi(game[0])
			if err != nil {
				log.Fatalf("Game count isn't a valid number: %s", string(group[0]))
			}
			color := game[1]
			groupData[color] += amount
		}
		for color, amount := range groupData {
			if gameCount[color] == 0 || gameCount[color] < amount {
				gameCount[color] = amount
			}
			if amount > gameConstraint[color] {
				possible = false
			}
		}
	}
	powerGroup := 1
	for _, amount := range gameCount {
		powerGroup *= amount
	}
	if possible {
		return gameID, powerGroup
	} else {
		return 0, powerGroup
	}
}
