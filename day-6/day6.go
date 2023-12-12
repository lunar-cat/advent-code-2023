package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const PartTwo = true

func main() {
	boatOptions()
}

func boatOptions() {
	file, err := os.Open("./day-6/input.txt")
	if err != nil {
		log.Fatal("Missing File")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	r := regexp.MustCompile(`\d+`)
	r2 := regexp.MustCompile(`\s+`)

	var times []int
	var distance []int
	initialSpeed := 0

	for scanner.Scan() {
		line := scanner.Text()
		if PartTwo {
			line = r2.ReplaceAllString(line, "")
		}
		if strings.HasPrefix(line, "Time:") {
			getData(line, &times, r)
		}
		if strings.HasPrefix(line, "Distance:") {
			getData(line, &distance, r)
		}
	}

	numOfWays := calculateWays(&times, &distance, initialSpeed)
	fmt.Println(numOfWays)
}

func calculateWays(time *[]int, distance *[]int, initialSpeed int) int {
	numberOfWays := 1
	for i := 0; i < len(*time); i++ {
		options := getOptions((*time)[i], (*distance)[i], initialSpeed)
		numberOfWays *= len(options)
	}
	return numberOfWays
}

func getData(line string, bucket *[]int, r *regexp.Regexp) {
	nums := r.FindAllString(line, -1)
	for _, n := range nums {
		mappedNum, _ := strconv.Atoi(n)
		*bucket = append(*bucket, mappedNum)
	}
}

func getOptions(time int, distance int, initialSpeed int) []int {
	var options []int
	for ms := 0; ms < time; ms++ {
		timeLeft := time - ms
		if initialSpeed*timeLeft > distance {
			options = append(options, ms)
		}
		initialSpeed++
	}
	return options
}
