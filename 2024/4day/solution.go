package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const pattern = "XMAS"

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

func parseInput(content []byte) []string {
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func partOneSolution(lines []string) int {
	count := 0
	for row := 0; row < len(lines); row++ {
		for column := 0; column < len(lines[row]); column++ {
			count += checkForwards(lines, row, column) +
				checkBackwards(lines, row, column) +
				checkUpwards(lines, row, column) +
				checkDownwards(lines, row, column) +
				checkUpRightDiagonal(lines, row, column) +
				checkDownRightDiagonal(lines, row, column) +
				checkDownLeftDiagonal(lines, row, column) +
				checkUpLeftDiagonal(lines, row, column)
		}
	}
	return count
}

func checkUpLeftDiagonal(lines []string, row int, column int) int {
	if row < 3 || column < 3 {
		return 0
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row-i][column-i])
	}

	if str == pattern {
		fmt.Println("UP LEFT DIAGONAL - ROW: ", row, "COLUMN: ", column)
		return 1
	}

	return 0
}

func checkDownLeftDiagonal(lines []string, row int, column int) int {
	if row > len(lines)-len(pattern) || column < 3 {
		return 0
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row+i][column-i])
	}

	if str == pattern {
		fmt.Println("DOWN LEFT DIAGONAL - ROW: ", row, "COLUMN: ", column)
		return 1
	}

	return 0
}

func checkDownRightDiagonal(lines []string, row int, column int) int {
	if row > len(lines)-len(pattern) || len(lines[row])-column < 4 {
		return 0
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row+i][column+i])
	}

	if str == pattern {
		fmt.Println("DOWN RIGHT DIAGONAL - ROW: ", row, "COLUMN: ", column)
		return 1
	}

	return 0
}

func checkUpRightDiagonal(lines []string, row int, column int) int {
	if row < 3 || len(lines[row])-column < 4 {
		return 0
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row-i][column+i])
	}

	if str == pattern {
		fmt.Println("UP RIGHT DIAGONAL - ROW: ", row, "COLUMN: ", column)
		return 1
	}

	return 0
}

func checkDownwards(lines []string, row int, column int) int {
	if row > len(lines)-len(pattern) {
		return 0
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row+i][column])
	}

	if str == pattern {
		fmt.Println("DOWNWARDS - ROW: ", row, "COLUMN: ", column)
		return 1
	}

	return 0
}

func checkUpwards(lines []string, row int, column int) int {
	if row < 3 {
		return 0
	}

	str := ""

	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row-i][column])
	}

	if str == pattern {
		fmt.Println("UPWARDS   - ROW: ", row, "COLUMN: ", column)
		return 1
	}

	return 0
}

func checkBackwards(lines []string, row int, column int) int {
	if column < 3 {
		return 0
	}

	str := ""
	i := 0

	for ; i < len(pattern); i++ {
		str += string(lines[row][column-i])
	}

	if str == pattern {
		fmt.Println("BACKWARDS - ROW: ", row, "COLUMN: ", column)
		return 1
	}

	return 0
}

func checkForwards(lines []string, row int, column int) int {
	if len(lines[row])-column < 4 {
		return 0
	}

	str := ""
	i := 0
	for ; i < len(pattern); i++ {
		str += string(lines[row][column+i])
	}

	if str == pattern {
		fmt.Println("FORWARDS  - ROW: ", row, "COLUMN: ", column)
		return 1
	}

	return 0
}

func partTwoSolution(inp []string) int {
	count := 0
	for row := 1; row < len(inp) - 1; row++ {
		for col := 1; col < len(inp[row]) - 1; col++ {
			firstDiagonal := string([]byte{inp[row-1][col-1], inp[row][col], inp[row+1][col+1]})
			secondDiagonal := string([]byte{inp[row+1][col-1], inp[row][col], inp[row-1][col+1]})

			if (firstDiagonal == "SAM" || firstDiagonal == "MAS") && (secondDiagonal == "SAM" || secondDiagonal == "MAS") {
				count++
			}

		}
	}

	return count
}
