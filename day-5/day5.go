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

func main() {
	almanac()
}

func almanac() {
	file, err := os.Open("./day-5/input.txt")
	if err != nil {
		log.Fatal("Missing File")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var seeds []int
	currentStage := -1
	var stages [7]map[[2]int]int

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "seeds") {
			seeds = getSeeds(line)
			continue
		}
		mapRanges(line, &stages, currentStage)
		currentStage = selectStage(line, currentStage)
	}
	lowestLocation := processSeeds(seeds, stages)
	fmt.Println(lowestLocation)
}

func getSeeds(line string) []int {
	r := regexp.MustCompile(`\d+`)
	seeds := r.FindAllString(line, -1)
	var mappedSeeds []int
	for _, seed := range seeds {
		numSeed, _ := strconv.Atoi(seed)
		mappedSeeds = append(mappedSeeds, numSeed)
	}
	return mappedSeeds
}

func selectStage(line string, currentStage int) int {
	switch {
	case strings.Contains(line, "seed-to-soil map"):
		return 0
	case strings.Contains(line, "soil-to-fertilizer map"):
		return 1
	case strings.Contains(line, "fertilizer-to-water map"):
		return 2
	case strings.Contains(line, "water-to-light map"):
		return 3
	case strings.Contains(line, "light-to-temperature map"):
		return 4
	case strings.Contains(line, "temperature-to-humidity map"):
		return 5
	case strings.Contains(line, "humidity-to-location map"):
		return 6
	}
	return currentStage
}

func mapRanges(line string, stagesMaps *[7]map[[2]int]int, currentStage int) {
	// destination - source - range (len)
	// 50 	   98 	   -> length 2
	// (98 - 98 + len - 1) : 50
	r := regexp.MustCompile(`\d+`)
	nums := r.FindAllString(line, -1)
	if len(nums) == 0 {
		return
	}
	var mappedNums [3]int
	for i, strNum := range nums {
		num, _ := strconv.Atoi(strNum)
		mappedNums[i] = num
	}
	rangeMap := stagesMaps[currentStage]
	k := [2]int{mappedNums[1], mappedNums[1] + (mappedNums[2] - 1)}
	v := mappedNums[0]
	if rangeMap == nil {
		rangeMap = map[[2]int]int{k: v}
	} else {
		rangeMap[k] = v
	}
	stagesMaps[currentStage] = rangeMap
}

func processSeeds(seeds []int, stagesMap [7]map[[2]int]int) int {
	var locations []int

	currentValue := 0
	for _, seed := range seeds { // seed -> 79
		currentValue = seed
	nextStage:
		for _, stage := range stagesMap { // stage -> {[1-10]: 12}
			for k, v := range stage { // k -> [1-10] | v -> 12
				if k[0] <= currentValue && currentValue <= k[1] {
					currentValue = (currentValue - k[0]) + v
					continue nextStage
				}
			}
		}
		locations = append(locations, currentValue)
	}
	lowestLocation := locations[0]

	for _, location := range locations {
		if location < lowestLocation {
			lowestLocation = location
		}
	}
	return lowestLocation
}
