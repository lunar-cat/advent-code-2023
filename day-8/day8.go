package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"
)

const PartTwo = true

func main() {
	start := time.Now()
	hauntedWasteland()
	fmt.Println(time.Since(start))
}

func hauntedWasteland() {
	file, err := os.Open("./day-8/input.txt")
	if err != nil {
		log.Fatal("Missing File")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var directions string
	nodes := make(map[string][2]string)
	startNode := "AAA"
	endNode := "ZZZ"

	for scanner.Scan() {
		line := scanner.Text()
		if directions == "" {
			directions = line
			continue
		}
		if line == "" {
			continue
		}
		listNodes(scanner.Text(), &nodes)
	}
	steps := 0
	if PartTwo {
		steps = traverseParallelNodes(&nodes, directions)
	} else {
		steps = traverseNodes(&nodes, startNode, endNode, directions)
	}
	fmt.Println(steps)
}

func listNodes(line string, nodes *map[string][2]string) {
	r := regexp.MustCompile(`[(),]+`)
	parts := strings.Split(line, " = ")
	node := parts[0]
	network := strings.Split(r.ReplaceAllString(parts[1], ""), " ")
	(*nodes)[node] = [2]string{network[0], network[1]}
}

func traverseNodes(nodes *map[string][2]string, startNode string, endNode string, directions string) int {
	steps := 0
	direction := 0
	currentNode := startNode
	for currentNode != endNode {
		nextMove := string(directions[(steps % len(directions))])
		steps++
		if nextMove == "R" {
			direction = 1
		} else {
			direction = 0
		}

		currentNode = (*nodes)[currentNode][direction]
	}
	return steps
}

func traverseParallelNodes(nodes *map[string][2]string, directions string) int {
	var currentNodes []string
	for k := range *nodes {
		if strings.HasSuffix(k, "A") {
			currentNodes = append(currentNodes, k)
		}
	}
	results := make([]int, len(currentNodes))
	var wg sync.WaitGroup

	for i, node := range currentNodes {
		wg.Add(1)
		go func(startNode string, index int) {
			defer wg.Done()
			steps := 0
			direction := 0
			currentNode := startNode
			for !strings.HasSuffix(currentNode, "Z") {
				nextMove := string(directions[(steps % len(directions))])
				steps++
				if nextMove == "R" {
					direction = 1
				} else {
					direction = 0
				}

				currentNode = (*nodes)[currentNode][direction]
			}
			results[index] = steps
		}(node, i)
	}
	wg.Wait()

	return LCM(results...)
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	result := integers[0] * integers[1] / GCD(integers[0], integers[1])

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
