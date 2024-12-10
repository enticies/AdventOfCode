package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

func parseInput(content []byte) [][]int {
	matrix := make([][]int, 0)

	for _, row := range strings.Split(strings.TrimSpace(string(content)), "\n") {
		rowChars := make([]int, 0)

		for _, char := range row {
			val, err := strconv.Atoi(string(char))

			if err != nil {
				val = -1
			}

			rowChars = append(rowChars, val)
		}
		matrix = append(matrix, rowChars)
	}

	return matrix
}

func findTrailHeads(matrix [][]int) []Coordinate {
	coordinates := make([]Coordinate, 0)

	for rowIndex := range matrix {
		for colIndex := range matrix[rowIndex] {
			if matrix[rowIndex][colIndex] == 0 {
				coordinates = append(coordinates, Coordinate{
					x: rowIndex,
					y: colIndex,
				})
			}
		}
	}

	return coordinates
}

var total = 0

func partOneSolution(matrix [][]int) int {
	trailHeads := findTrailHeads(matrix)

	total := 0

	for _, trailHead := range trailHeads {
		u := make(map[Coordinate]bool)

		walkRecursive(matrix, trailHead, u)

		total += len(u)
	}

	return total
}

func walkRecursivePartTwo(matrix [][]int, coordinate Coordinate) int {
	if matrix[coordinate.x][coordinate.y] == 9 {
		return 1
	}

	total := 0

	if isValidCoordinate(matrix, coordinate, "DOWN") && matrix[coordinate.x+1][coordinate.y]-matrix[coordinate.x][coordinate.y] == 1 {
		total += walkRecursivePartTwo(matrix, Coordinate{
			x: coordinate.x + 1,
			y: coordinate.y,
		})
	}
	if isValidCoordinate(matrix, coordinate, "RIGHT") && matrix[coordinate.x][coordinate.y+1]-matrix[coordinate.x][coordinate.y] == 1 {
		total += walkRecursivePartTwo(matrix, Coordinate{
			x: coordinate.x,
			y: coordinate.y + 1,
		})
	}

	if isValidCoordinate(matrix, coordinate, "LEFT") && matrix[coordinate.x][coordinate.y-1]-matrix[coordinate.x][coordinate.y] == 1 {
		total += walkRecursivePartTwo(matrix, Coordinate{
			x: coordinate.x,
			y: coordinate.y - 1,
		})
	}
	if isValidCoordinate(matrix, coordinate, "UP") && matrix[coordinate.x-1][coordinate.y]-matrix[coordinate.x][coordinate.y] == 1 {
		total += walkRecursivePartTwo(matrix, Coordinate{
			x: coordinate.x - 1,
			y: coordinate.y,
		})
	}

	return total 
}

func walkRecursive(matrix [][]int, coordinate Coordinate, u map[Coordinate]bool) {
	if matrix[coordinate.x][coordinate.y] == 9 {
		u[Coordinate{x: coordinate.x, y: coordinate.y}] = true
		return
	}

	if isValidCoordinate(matrix, coordinate, "DOWN") && matrix[coordinate.x+1][coordinate.y]-matrix[coordinate.x][coordinate.y] == 1 {
		walkRecursive(matrix, Coordinate{
			x: coordinate.x + 1,
			y: coordinate.y,
		}, u)
	}
	if isValidCoordinate(matrix, coordinate, "RIGHT") && matrix[coordinate.x][coordinate.y+1]-matrix[coordinate.x][coordinate.y] == 1 {
		walkRecursive(matrix, Coordinate{
			x: coordinate.x,
			y: coordinate.y + 1,
		}, u)
	}

	if isValidCoordinate(matrix, coordinate, "LEFT") && matrix[coordinate.x][coordinate.y-1]-matrix[coordinate.x][coordinate.y] == 1 {
		walkRecursive(matrix, Coordinate{
			x: coordinate.x,
			y: coordinate.y - 1,
		}, u)
	}
	if isValidCoordinate(matrix, coordinate, "UP") && matrix[coordinate.x-1][coordinate.y]-matrix[coordinate.x][coordinate.y] == 1 {
		walkRecursive(matrix, Coordinate{
			x: coordinate.x - 1,
			y: coordinate.y,
		}, u)
	}

	return
}

func isAtCorner(matrix [][]int, coordinate Coordinate) bool {
	return coordinate.x == 0 || coordinate.x == len(matrix)-1 || coordinate.y == 0 || coordinate.y == len(matrix[0])-1
}

func isValidCoordinate(matrix [][]int, coordinate Coordinate, direction string) bool {
	if direction == "UP" && coordinate.x > 0 {
		return true
	} else if direction == "RIGHT" && coordinate.y < len(matrix[0])-1 {
		return true
	} else if direction == "DOWN" && coordinate.x < len(matrix)-1 {
		return true
	} else if direction == "LEFT" && coordinate.y > 0 {
		return true
	}

	return false
}

type Coordinate struct {
	x int
	y int
}

func printMatrix(matrix [][]int) {
	maxWidth := 0
	for _, row := range matrix {
		for _, val := range row {
			width := 1
			if val != -1 {
				width = len(fmt.Sprint(val))
			}
			if width > maxWidth {
				maxWidth = width
			}
		}
	}

	for _, row := range matrix {
		for i, val := range row {
			if val == -1 {
				fmt.Print(strings.Repeat(" ", maxWidth-1) + ".")
			} else {
				fmt.Printf("%*d", maxWidth, val)
			}
			if i < len(row)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func partTwoSolution(matrix [][]int) int {
	trailHeads := findTrailHeads(matrix)

	total := 0

	for _, trailHead := range trailHeads {
		total += walkRecursivePartTwo(matrix, trailHead)
	}

	return total
}
