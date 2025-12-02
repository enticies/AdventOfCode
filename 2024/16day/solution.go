package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	filePath := "example_input.txt"

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	inp := parseInput(content)

	partOneResult := partOneSolution(inp)
	// partTwoResult := partTwoSolution(copyMatrix(inp))

	fmt.Printf("Part one result is %d\n", partOneResult)
	// fmt.Printf("Part two result is %d\n", partTwoResult)
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

type Coordinate struct {
	x, y int
}

var directions = map[string]Coordinate{
	"up":    {0, -1},
	"right": {1, 0},
	"down":  {0, 1},
	"left":  {1, 0},
}

func getNodes(matrix [][]string, coor Coordinate) []Coordinate {
	coordinates := make([]Coordinate, 0)

	for _, direction := range directions {
		newX := coor.x + direction.x
		newY := coor.y + direction.y

		if newX > 0 && newX < len(matrix) && newY > 0 && newY < len(matrix[0]) && matrix[newY][newX] != "#" {
			coordinates = append(coordinates, Coordinate{newX, newY})
		}
	}
	return coordinates
}

func printMatrix(lines [][]string) {
	for _, row := range lines {
		fmt.Println(row)
	}
	fmt.Println()
}

func partOneSolution(lines [][]string) int {
	return 0
}
