package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const pattern = "XMAS"

func main() {
	filePath := "example_input.txt"

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

func parseInput(content []byte) [][]string {
	splitLines := strings.Split(strings.TrimSpace(string(content)), "\n")

	result := make([][]string, 0)

	for _, line := range splitLines {
		result = append(result, strings.Split(line, ":"))
	}

	return result
}

func partOneSolution(lines [][]string) int {
	total := 0

	for _, line := range lines {
		testValue, _ := strconv.Atoi(line[0])

		values := make([]int, 0)

		for _, value := range strings.Split(line[1], " ") {
			convertedValue, _ := strconv.Atoi(value)

			values = append(values, convertedValue)
		}

		if canEqual(testValue, values[1:], values[0]) {
			total += testValue
			fmt.Println("CAN EQUAL")
		} else {
			fmt.Println("CAN'T EQUAL")
		}

		break
	}

	return total
}

var operations = []string{"ADD", "MULTIPLY", "CONCATENATION"}

func canEqual(testValue int, values []int, currentTotal int) bool {
	fmt.Println("TEST VALUE: ", testValue)
	fmt.Println("VALUES: ", values)
	fmt.Println("CURRENT TOTAL: ", currentTotal)
	fmt.Println()
	if len(values) == 0 {
		return testValue == currentTotal
	}

	result := false

	for _, operation := range operations {
		if canEqual(testValue, values[1:], executeOperation(currentTotal, values[0], operation)) {
			return true
		}
	}

	return result
}

func executeOperation(currentTotal int, value int, operation string) int {
	if operation == "ADD" {
		return currentTotal + value
	} else if operation == "MULTIPLY" {
		return currentTotal * value
	} else if operation == "CONCATENATION" {
		concatenatedInts, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentTotal, value))
		return concatenatedInts 
	} 

	return currentTotal
}

func partTwoSolution(lines [][]string) int {
	total := 0

	for _, line := range lines {
		testValue, _ := strconv.Atoi(line[0])

		values := make([]int, 0)

		for _, value := range strings.Split(line[1], " ") {
			convertedValue, _ := strconv.Atoi(value)

			values = append(values, convertedValue)
		}

		if canEqual(testValue, values[1:], values[0]) {
			total += testValue
			fmt.Println("CAN EQUAL")
		} else {
			fmt.Println("CAN'T EQUAL")
		}
	}

	return total
}
