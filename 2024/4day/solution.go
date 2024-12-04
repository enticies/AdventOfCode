package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	// partTwoResult := partTwoSolution(inp)

	fmt.Printf("Part one result is %d\n", partOneResult)
	// fmt.Printf("Part two result is %d\n", partTwoResult)
}

func parseInput(content []byte) []string {
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func partOneSolution(lines []string) int {
	count := 0
	for row := 0; row < len(lines); row++ {
		for column := 0; column < len(lines[row]) - 1; column++ {
			if checkForwards(lines, row, column) || 
			   checkBackwards(lines, row, column) || 
			   checkUpwards(lines, row, column) || 
			   checkDownwards(lines, row, column) || 
			   checkUpRightDiagonal(lines, row, column) || 
			   checkDownRightDiagonal(lines, row, column) || 
			   checkDownLeftDiagonal(lines, row, column) ||
			   checkUpLeftDiagonal(lines, row, column) {
				count++
			}
		}
	}
	return count
}

func checkUpLeftDiagonal(lines []string, row int, column int) bool {
	if row < 3 || column < 3 {
		return false
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row-i][column-i])
	}

	if str == pattern {
		fmt.Println("UP LEFT DIAGONAL - ROW: ", row, "COLUMN: ", column)
		return true
	}

	return false
}

func checkDownLeftDiagonal(lines []string, row int, column int) bool {
	if row > len(lines)-len(pattern)  || column < 3 {
		return false
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row+i][column-i])
	}

	if str == pattern {
		fmt.Println("DOWN LEFT DIAGONAL - ROW: ", row, "COLUMN: ", column)
		return true
	}

	return false
}

func checkDownRightDiagonal(lines []string, row int, column int) bool {
	if row > len(lines)-len(pattern) || len(lines[row])-column - 1 < 4 {
		return false
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row+i][column+i])
	}

	if str == pattern {
		fmt.Println("DOWN RIGHT DIAGONAL - ROW: ", row, "COLUMN: ", column)
		return true
	}

	return false
}

func checkUpRightDiagonal(lines []string, row int, column int) bool {
	if row < 3 || len(lines[row])-column < 4 {
		return false
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row-i][column+i])
	}

	if str == pattern {
		fmt.Println("UP RIGHT DIAGONAL - ROW: ", row, "COLUMN: ", column)
		return true
	}

	return false
}

func checkDownwards(lines []string, row int, column int) bool {
	if row > len(lines)-len(pattern) {
		return false
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row+i][column])
	}

	if str == pattern {
		fmt.Println("DOWNWARDS - ROW: ", row, "COLUMN: ", column)
		return true
	}

	return false
}

func checkUpwards(lines []string, row int, column int) bool {
	if row < 3 {
		return false
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row-i][column])
	}

	if str == pattern {
		fmt.Println("UPWARDS   - ROW: ", row, "COLUMN: ", column)
		return true
	}

	return false
}

func checkBackwards(lines []string, row int, column int) bool {
	if column < 3 {
		return false
	}

	str := ""
	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row][column-i])
	}

	if str == pattern {
		fmt.Println("BACKWARDS - ROW: ", row, "COLUMN: ", column)
		return true
	}

	return false
}

func checkForwards(lines []string, row int, column int) bool {
	if len(lines[row])-column < 4 {
		return false
	}

	str := ""
	i := 0
	for ; i < len(pattern); i++ {
		str += string(lines[row][column+i])
	}

	if str == pattern {
		fmt.Println("FORWARDS  - ROW: ", row, "COLUMN: ", column)
		return true
	}

	return false
}

func partTwoSolution(inp string) int {
	return 0
}
