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
	almanac()
}

func almanac() {
	file, err := os.Open("./day-5/input.txt")
	if err != nil {
		log.Fatal("Missing File")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var seeds [][2]int
	currentStage := -1
	var stages [7]map[[2]int]int
	r := regexp.MustCompile(`\d+`)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "seeds") {
			seeds = getSeeds(line, r)
			continue
		}
		mapRanges(line, &stages, currentStage)
		currentStage = selectStage(line, currentStage)
	}
	lowestLocation := processSeeds(seeds, stages)
	fmt.Println(lowestLocation)
}

func getSeeds(line string, r *regexp.Regexp) [][2]int {
	seeds := r.FindAllString(line, -1)
	var mappedSeeds [][2]int

	if !PartTwo {
		for i := 0; i < len(seeds); i++ {
			startSeed, _ := strconv.Atoi(seeds[i])
			mappedSeeds = append(mappedSeeds, [2]int{startSeed, startSeed + 1})
		}
	} else {
		for i := 0; i < len(seeds); i += 2 {
			startSeed, _ := strconv.Atoi(seeds[i])
			endSeed, _ := strconv.Atoi(seeds[i+1])
			mappedSeeds = append(mappedSeeds, [2]int{startSeed, startSeed + endSeed})
		}
	}
	return mappedSeeds
}

func selectStage(line string, currentStage int) int {
	switch {
	case strings.HasPrefix(line, "seed-to-soil"):
		return 0
	case strings.HasPrefix(line, "soil-to-fertilizer"):
		return 1
	case strings.HasPrefix(line, "fertilizer-to-water"):
		return 2
	case strings.HasPrefix(line, "water-to-light"):
		return 3
	case strings.HasPrefix(line, "light-to-temperature"):
		return 4
	case strings.HasPrefix(line, "temperature-to-humidity"):
		return 5
	case strings.HasPrefix(line, "humidity-to-location"):
		return 6
	}
	return currentStage
}

func mapRanges(line string, stagesMaps *[7]map[[2]int]int, currentStage int) {
	// destination - source - range (len)
	// 50 	   98 	   -> length 2
	// (98 - 98 + len) : 50
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
	k := [2]int{mappedNums[1], mappedNums[1] + (mappedNums[2])}
	v := mappedNums[0]
	if rangeMap == nil {
		rangeMap = map[[2]int]int{k: v}
	} else {
		rangeMap[k] = v
	}
	stagesMaps[currentStage] = rangeMap
}

func processSeeds(seeds [][2]int, stagesMap [7]map[[2]int]int) int {
	var locations [][2]int
	var newSeeds [][2]int

	for _, stage := range stagesMap {
		newSeeds = [][2]int{}
		for len(seeds) > 0 {
			i := len(seeds) - 1
			seedRange := seeds[i]
			seeds = append(seeds[:i], seeds[i+1:]...)
			found := false
			for source, destination := range stage { // source [start, end] : destination
				lowerIntersection := max(seedRange[0], source[0])
				upperIntersection := min(seedRange[1], source[1])
				if lowerIntersection < upperIntersection {
					newSeeds = append(newSeeds, [2]int{ // diff -> source start - destination
						lowerIntersection - source[0] + destination,
						upperIntersection - source[0] + destination,
					})
					if seedRange[0] < lowerIntersection {
						seeds = append(seeds, [2]int{seedRange[0], lowerIntersection})
					}
					if seedRange[1] > upperIntersection {
						seeds = append(seeds, [2]int{upperIntersection, seedRange[1]})
					}
					found = true
				}
			}
			if !found {
				newSeeds = append(newSeeds, seedRange)
			}
		}
		seeds = newSeeds
	}
	locations = seeds

	lowestLocation := locations[0]

	for _, location := range locations {
		if location[0] < lowestLocation[0] {
			lowestLocation = location
		}
	}
	return lowestLocation[0]
}
