package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode"
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

func parseInput(content []byte) [][]rune {
	splitLines := strings.Split(strings.TrimSpace(string(content)), "\n")

	matrix := make([][]rune, 0)

	for _, line := range splitLines {
		row := make([]rune, 0)

		for _, char := range line {
			row = append(row, char)
		}

		matrix = append(matrix, row)
	}

	return matrix
}

func isAntenna(char rune) bool {
	return unicode.IsDigit(char) ||
		unicode.IsUpper(char) ||
		unicode.IsLower(char)
}

type Coordinate struct {
	x int
	y int
}

func partOneSolution(matrix [][]rune) int {
	total := 0

	antennas := getAntennaLocations(matrix)

	for _, value := range antennas {
		lines := getPermutationsOfTwo(value)

		for _, line := range lines {
			diffX, diffY := calculateDifferenceOfCoordinate(line[0], line[1])
			if antinodeExists(matrix, line[1], diffX, diffY) {
				total++
			}
		}
	}

	return total
}

func antinodeExists(matrix [][]rune, currentLocation Coordinate, diffX int, diffY int) bool {
	newX := currentLocation.x + diffX
	newY := currentLocation.y + diffY

	if newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) {
		if matrix[newX][newY] == '#' {
			return false
		}
		matrix[newX][newY] = '#'
		return true
	}

	return false
}

func findAntiNodes(matrix [][]rune, currentLocation Coordinate, diffX int, diffY int) int {
	total := 0

	newX := currentLocation.x + diffX
	newY := currentLocation.y + diffY

	for newX >= 0 && newX < len(matrix) && newY >= 0 && newY < len(matrix[0]) {
		if !isAntenna(matrix[newX][newY]) && matrix[newX][newY] != '#' {
			matrix[newX][newY] = '#'
			total++
		}

		newX += diffX
		newY += diffY
	}

	return total
}

func calculateDifferenceOfCoordinate(coorOne Coordinate, coorTwo Coordinate) (int, int) {
	return coorTwo.x - coorOne.x, coorTwo.y - coorOne.y
}

func getPermutationsOfTwo(coords []Coordinate) [][2]Coordinate {
	result := make([][2]Coordinate, 0)

	if len(coords) < 2 {
		return result
	}

	for i := 0; i < len(coords); i++ {
		for j := 0; j < len(coords); j++ {
			if i == j {
				continue
			}
			result = append(result, [2]Coordinate{coords[i], coords[j]})
		}
	}

	return result
}

func getAntennaLocations(matrix [][]rune) map[rune][]Coordinate {
	antennas := make(map[rune][]Coordinate)

	for ri := range matrix {
		for ci := range matrix {
			if isAntenna(matrix[ri][ci]) {
				if _, exists := antennas[matrix[ri][ci]]; exists {
					antennas[matrix[ri][ci]] = append(antennas[matrix[ri][ci]], Coordinate{x: ri, y: ci})
				} else {
					antennas[matrix[ri][ci]] = make([]Coordinate, 0)
					antennas[matrix[ri][ci]] = append(antennas[matrix[ri][ci]], Coordinate{x: ri, y: ci})
				}
			}
		}
	}

	return antennas
}

func printMatrix(matrix [][]rune) {
	for i := range matrix[0] {
		if i > 9 {
			fmt.Print(fmt.Sprintf("%x", i))
		} else {
			fmt.Print(i)
		}
	}

	fmt.Println()

	for ri, row := range matrix {
		for _, col := range row {
			fmt.Printf("%c", col)
		}
		fmt.Printf(" %d\n", ri)
	}

	for i := range matrix[0] {
		if i > 9 {
			fmt.Print(fmt.Sprintf("%x", i))
		} else {
			fmt.Print(i)
		}
	}

}

func partTwoSolution(matrix [][]rune) int {
	total := 0

	antennas := getAntennaLocations(matrix)

	for _, value := range antennas {
		lines := getPermutationsOfTwo(value)

		for _, line := range lines {
			diffX, diffY := calculateDifferenceOfCoordinate(line[0], line[1])
			total += findAntiNodes(matrix, line[1], diffX, diffY)
		}
	}

	for _, value := range antennas {
		total += len(value)
	}

	return total
}
