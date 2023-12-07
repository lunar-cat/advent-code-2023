package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type NumberData struct {
	value string
	start int
	end   int
	valid bool
	clear bool
}

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
	sumID := getSchematicValues(scanner)
	fmt.Println(sumID)
}

func getSchematicValues(scanner *bufio.Scanner) int {
	sum := 0
	num := 0
	pLine := ""
	for scanner.Scan() {
		pLine, num = getAdjacentNumbers(pLine, scanner.Text())
		sum += num
	}
	return sum
}

func getAdjacentNumbers(pLine string, cLine string) (string, int) {
	var numbers []NumberData
	var validNumbers []int
	var marksIndex []int
	addingNumber := false
	clearValue := false
	numData := NumberData{
		value: "",
		start: 0,
		end:   0,
		valid: false,
		clear: false,
	}
	line := ""
	for i := 0; i < 2; i++ {
		if i == 0 {
			line = pLine
		} else {
			clearValue = true
			line = cLine
		}
		for i, char := range line {
			if unicode.IsDigit(char) {
				if addingNumber == false {
					clearStruct(&numData)
					addingNumber = true
					numData.start = i
				}
				numData.value += string(char)
				if i+1 == len(line) {
					numData.end = i - 1
					numData.clear = clearValue
					addingNumber = false
					numbers = append(numbers, numData)
				}
			} else {
				if addingNumber {
					numData.end = i - 1
					numData.clear = clearValue
					addingNumber = false
					numbers = append(numbers, numData)
				}
			}
			if !unicode.IsDigit(char) && char != '.' {
				marksIndex = append(marksIndex, i)
			}
		}
	}

	for _, mark := range marksIndex {
		for i, number := range numbers {
			if (number.start-1) <= mark && (number.end+1) >= mark && !number.valid {
				numbers[i].valid = true
				if number.clear {
					cLine = replaceSubstring(cLine, number.start, number.end, strings.Repeat(".", len(number.value)))
				}
			}
		}
	}

	sum := 0
	for _, number := range numbers {
		if !number.valid {
			continue
		}
		num, err := strconv.Atoi(number.value)
		validNumbers = append(validNumbers, num)
		if err == nil {
			sum += num
		}
	}
	return cLine, sum
}

func clearStruct(v interface{}) {
	p := reflect.ValueOf(v).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func replaceSubstring(original string, start, end int, newSubstring string) string {
	var builder strings.Builder

	builder.WriteString(original[:start])
	builder.WriteString(newSubstring)
	builder.WriteString(original[end+1:])
	return builder.String()
}
