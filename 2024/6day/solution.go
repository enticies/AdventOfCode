package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

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
	spl := strings.Split(strings.TrimSpace(string(content)), "\n")

	lines := make([][]string, 0)

	for _, line := range spl {
		chars := make([]string, 0)

		for _, char := range line {
			chars = append(chars, string(char))
		}

		lines = append(lines, chars)
	}

	return lines
}

func partOneSolution(lines [][]string) int {
	row, col := findGuardPosition(lines)

	for row != 0 && row != len(lines)-1 && col != 0 && col != len(lines[0])-1 {
		guard := lines[row][col]
		direction := getDirection(guard)

		if direction == "UP" {
			if lines[row-1][col] == "#" {
				lines[row][col] = ">"
			} else {
				lines[row][col] = "%"
				lines[row-1][col] = "^"
			}
		} else if direction == "RIGHT" {
			if lines[row][col+1] == "#" {
				lines[row][col] = "v"
			} else {
				lines[row][col] = "%"
				lines[row][col+1] = ">"
			}
		} else if direction == "DOWN" {
			if lines[row+1][col] == "#" {
				lines[row][col] = "<"
			} else {
				lines[row][col] = "%"
				lines[row+1][col] = "v"
			}
		} else if direction == "LEFT" {
			if lines[row][col-1] == "#" {
				lines[row][col] = "^"
			} else {
				lines[row][col] = "%"
				lines[row][col-1] = "<"
			}
		}

		printMatrix(lines)
		row, col = findGuardPosition(lines)
	}

	return countVisitedPosition(lines) + 1
}

func countVisitedPosition(lines [][]string) int {
	total := 0

	for _, line := range lines {
		for _, char := range line {
			if char == "%" {
				total++
			}
		}
	}

	return total
}

func printMatrix(lines [][]string) {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	for _, row := range lines {
		fmt.Println(row)
	}
	fmt.Println()
	time.Sleep(50 * time.Millisecond)
}

func findGuardPosition(lines [][]string) (row int, column int) {
	for row, line := range lines {
		for column, char := range line {
			if char == "^" || char == ">" || char == "<" || char == "v" {
				return row, column
			}
		}
	}
	return
}

func getDirection(char string) string {
	ret := ""
	if char == "^" {
		ret = "UP"
	} else if char == ">" {
		ret = "RIGHT"
	} else if char == "v" {
		ret = "DOWN"
	} else if char == "<" {
		ret = "LEFT"
	}
	return ret
}

type Pair struct {
	First  int
	Second int
}

func createsLoop(lines [][]string, row int, col int) bool {
	visitedPositions := make(map[Pair]string)
	loop := false


	for row != 0 && row != len(lines)-1 && col != 0 && col != len(lines[0])-1 {
		guard := lines[row][col]
		direction := getDirection(guard)

		if direction == "UP" {
			if val, exists := visitedPositions[Pair{row, col}]; exists {
				if val == "UP" {
					loop = true
					break
				}
			} else {
				visitedPositions[Pair{row, col}] = "UP"
			}

			if lines[row-1][col] == "#" {
				lines[row][col] = ">"
			} else {
				lines[row][col] = "%"
				lines[row-1][col] = "^"
				row--
			}

		} else if direction == "RIGHT" {
			if val, exists := visitedPositions[Pair{row, col}]; exists {
				if val == "RIGHT" {
					loop = true
					break
				}
			} else {
				visitedPositions[Pair{row, col}] = "RIGHT"
			}

			if lines[row][col+1] == "#" {
				lines[row][col] = "v"
			} else {
				lines[row][col] = "%"
				lines[row][col+1] = ">"
				col++
			}

		} else if direction == "DOWN" {
			if val, exists := visitedPositions[Pair{row, col}]; exists {
				if val == "DOWN" {
					loop = true
					break
				}
			} else {
				visitedPositions[Pair{row, col}] = "DOWN"
			}

			if lines[row+1][col] == "#" {
				lines[row][col] = "<"
			} else {
				lines[row][col] = "%"
				lines[row+1][col] = "v"
				row++
			}

		} else if direction == "LEFT" {
			if val, exists := visitedPositions[Pair{row, col}]; exists {
				if val == "LEFT" {
					loop = true
					break
				}
			} else {
				visitedPositions[Pair{row, col}] = "LEFT"
			}

			if lines[row][col-1] == "#" {
				lines[row][col] = "^"
			} else {
				lines[row][col] = "%"
				lines[row][col-1] = "<"
				col--
			}
		}

		printMatrix(lines)
	}

	return loop
}

func copyMatrix(grid [][]string) [][]string {
    new_copy := make([][]string, len(grid))
    for i := range grid {
        new_copy[i] = make([]string, len(grid[i]))
        copy(new_copy[i], grid[i]) 
    }
    return new_copy
}

func partTwoSolution(lines [][]string) int {
	total := 0

	row, col := findGuardPosition(lines)

	for r := range lines {
		for c := range lines[r] {
			if lines[r][c] == "#" || lines[r][c] == "^" || lines [r][c] == ">" || lines[r][c] == "v" || lines[r][c] == "<" {
				continue
			}

			new_copy := copyMatrix(lines)
			new_copy[r][c] = "#"

			if createsLoop(new_copy, row, col) {
				total++
			}
		}
	}

	return total
}
