package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	calibration()
}

func calibration() {
	sum := 0
	file, err := os.Open("./day-1/input.txt")
	if err != nil {
		log.Fatal("Missing File")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := getNumberFromLine(scanner.Text())
		if err == nil {
			sum += num
		}
	}
	fmt.Println(sum)
}

func getNumberFromLine(line string) (int, error) {
	replaceMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	for k, v := range replaceMap {
		line = strings.ReplaceAll(line, k, string(k[0])+v+string(k[len(k)-1]))
	}

	var number string
	for _, char := range line {
		if unicode.IsDigit(char) {
			number += string(char)
			continue
		}
	}

	firstDigit, _ := strconv.Atoi(string(number[0]))
	lastDigit, _ := strconv.Atoi(string(number[len(number)-1]))
	return firstDigit*10 + lastDigit, nil
}
