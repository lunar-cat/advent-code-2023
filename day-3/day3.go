package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	engine()
}

func engine() {
	file, err := os.Open("./day-3/input.txt")
	if err != nil {
		log.Fatal("Missing File")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}
	sum, pow := getSchematicValues(data)
	fmt.Println(sum, pow)
}

func getSchematicValues(data []string) (int, int) {
	totalSum := 0
	totalPow := 0
	symbolsPosition := make(map[[2]int][]int)
	for i, line := range data {
		getSymbolsPositions(line, symbolsPosition, i)
	}
	for i, line := range data {
		getTouchingNumbers(line, symbolsPosition, i)
	}
	for _, values := range symbolsPosition {
		totalSum += getSum(values)
		if len(values) == 2 {
			totalPow += values[0] * values[1]
		}
	}
	return totalSum, totalPow
}

func getSymbolsPositions(line string, symbolsMap map[[2]int][]int, row int) {
	r := regexp.MustCompile("[^\\d.]")
	foundIndex := r.FindAllStringIndex(line, -1) // n is the limit of match (-1 no limit)
	for _, pos := range foundIndex {
		symbolsMap[[2]int{row, pos[0]}] = []int{}
	}
}

func getTouchingNumbers(line string, symbolsMap map[[2]int][]int, row int) {
	r := regexp.MustCompile("\\d+")
	foundIndex := r.FindAllStringIndex(line, -1)
numLoop:
	for _, pos := range foundIndex {
		colStart := pos[0] - 1
		colEnd := pos[1]
		rowStart := row - 1
		rowEnd := row + 1
		num, _ := strconv.Atoi(line[pos[0]:pos[1]])
		for key := range symbolsMap {
			keyRow := key[0]
			keyCol := key[1]
			if (keyCol >= colStart && keyCol <= colEnd) && (keyRow >= rowStart && keyRow <= rowEnd) {
				prevValues := symbolsMap[[2]int{keyRow, keyCol}]
				updatedValues := append(prevValues, num)
				symbolsMap[[2]int{keyRow, keyCol}] = updatedValues
				continue numLoop
			}
		}
	}
}

func getSum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
