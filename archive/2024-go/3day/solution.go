package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	inp := parseInput(content)

	partOneResult := partOneSolution(inp)
	partTwoResult := partTwoSolution(inp)

	fmt.Printf("Part one result is %d\n", partOneResult)
	fmt.Printf("Part two result is %d\n", partTwoResult)
}

func parseInput(content []byte) string {
	return strings.TrimSpace(string(content))
}

func partOneSolution(inp string) int {
	pattern := `\bmul\(\s*(-?\d+)\s*,\s*(-?\d+)\s*\)`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllSubmatch([]byte(inp), -1)
	total := 0

	for _, match := range matches {
		firstNum, _ := strconv.Atoi(string(match[1]))
		secondNum, _ := strconv.Atoi(string(match[2]))

		total += firstNum * secondNum
	}

	return total
}

func findLastDoIndex(text string) int {
	pattern := `do\(\)`
	regex := regexp.MustCompile(pattern)
	match := regex.FindStringIndex(text)

	if match == nil {
		return -1
	}
	return match[1]
}

func partTwoSolution(inp string) int {
	total := 0

	doNots := strings.Split(inp, "don't()")

	for i, val := range doNots {
		lastDoIndex := findLastDoIndex(val)
		pattern := `mul\(\s*(-?\d+)\s*,\s*(-?\d+)\s*\)`
		regex := regexp.MustCompile(pattern)
		substring := val

		if i == 0 {
			substring = val
		} else if (lastDoIndex == -1) {
			continue
		} else {
			substring = val[lastDoIndex:]
		}


		matches := regex.FindAllSubmatch([]byte(substring), -1)

		for _, match := range matches {
			firstNum, _ := strconv.Atoi(string(match[1]))
			secondNum, _ := strconv.Atoi(string(match[2]))

			total += firstNum * secondNum
		}
	}

	return total
}
