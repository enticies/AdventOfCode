package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	rows := parseInput(content)

	partOneResult := partOneSolution(rows)
	partTwoResult := partTwoSolution(rows)

	fmt.Printf("Part one result is %d\n", partOneResult)
	fmt.Printf("Part two result is %d\n", partTwoResult)
}

func parseInput(content []byte) [][]int {
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	rows := make([][]int, 0)

	for _, line := range lines {
		splitLine := strings.Fields(line)

		numberLine := make([]int, 0)

		for _, val := range splitLine {
			num, err := strconv.Atoi(val)

			if err != nil {
				panic(err)
			}

			numberLine = append(numberLine, num)
		}

		rows = append(rows, numberLine)
	}

	return rows
}

func partOneSolution(rows [][]int) int {
	unsafeLevels := 0
	for _, row := range rows {
		if !isRowSafe(row) {
			unsafeLevels++
		}
	}
	return len(rows) - unsafeLevels
}

func partTwoSolution(rows [][]int) int {
	safeLevels := 0
	for _, row := range rows {
		if isRowSafe(row) {
			safeLevels++
			fmt.Println(row, "- SAFE")
		} else if isSafeByRemoving(row) {
			safeLevels++
		} else {
			fmt.Println(row, "- UNSAFE")
		}
	}
	return safeLevels
}

func isSafeByRemoving(row []int) bool {
	for j := 0; j < len(row); j++ {
		newRow := make([]int, 0, len(row)-1)
		newRow = append(newRow, row[:j]...)
		newRow = append(newRow, row[j+1:]...)

		if isRowSafe(newRow) {
			fmt.Println(row, " SAFE BY REMOVING ", row[j])
			return true
		}
	}
	return false
}

func isRowSafe(row []int) bool {
	ascending := false

	if row[0] < row[1] {
		ascending = true
	}

	for i := 1; i < len(row); i++ {
		if (row[i-1] <= row[i] && !ascending) || (row[i-1] >= row[i] && ascending) || math.Abs(float64(row[i-1])-float64(row[i])) > 3 || math.Abs(float64(row[i-1])-float64(row[i])) < 1 {
			return false
		}
	}
	return true
}
