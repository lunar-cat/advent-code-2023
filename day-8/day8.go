package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

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
	steps := traverseNodes(&nodes, startNode, endNode, directions)
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
		if currentNode == endNode {
			break
		}
	}
	return steps
}
