package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"unicode"
)

func main() {
	filePath := "example_input.txt"

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	inp := parseInput(content)

	// partOneResult := partOneSolution(copyMatrix(inp))
	partTwoResult := partTwoSolution(copyMatrix(inp))

	// fmt.Printf("Part one result is %d\n", partOneResult)
	fmt.Printf("Part two result is %d\n", partTwoResult)
}

func copyMatrix(grid [][]string) [][]string {
    new_copy := make([][]string, len(grid))
    for i := range grid {
        new_copy[i] = make([]string, len(grid[i]))
        copy(new_copy[i], grid[i]) 
    }
    return new_copy
}

var directions = map[string]Coordinate{
	"up":    {-1, 0},
	"right": {0, 1},
	"down":  {1, 0},
	"left":  {0, -1},
}

type Coordinate struct {
	x, y int
}

func parseInput(content []byte) [][]string {
	s := make([][]string, 0)

	splitContent := strings.Split(strings.TrimSpace(string(content)), "\n")

	for _, row := range splitContent {
		s = append(s, strings.Split(row, ""))
	}

	return s
}

func IsSingleLetterLowercase(s string) bool {
	return unicode.IsLower(rune(s[0]))
}

func partOneSolution(matrix [][]string) int {
	total := 0

	for x := range matrix {
		for y := range matrix[x] {
			if !IsSingleLetterLowercase(matrix[x][y]) {
				area, perimeter := calculateRegion(matrix, Coordinate{x, y})
				total += area * perimeter
			}
		}
	}

	return total
}

func calculateRegion(matrix [][]string, startingCoordinate Coordinate) (int, int) {
	area := 0
	perimeter := 0

	unvisitedCoordinates := make([]Coordinate, 0)

	unvisitedCoordinates = append(unvisitedCoordinates, startingCoordinate)

	currentChar := matrix[startingCoordinate.x][startingCoordinate.y]

	for len(unvisitedCoordinates) > 0 {
		for _, coor := range unvisitedCoordinates {
			perimeter += calculatePlotPerimeter(matrix, coor, currentChar)
			for _, c := range unvisitedCoordinates {
				matrix[c.x][c.y] = strings.ToLower(matrix[c.x][c.y])
			}
			area++
			unvisitedCoordinates = unvisitedCoordinates[1:]
			unvisitedCoordinates = append(unvisitedCoordinates, getUnvisitedNodes(matrix, coor, currentChar)...)
		}
	}

	return area, perimeter
}

func calculatePlotPerimeter(matrix [][]string, coor Coordinate, char string) int {
	total := 0

	for _, direction := range directions {
		newX := coor.x + direction.x
		newY := coor.y + direction.y

		if newX < 0 || newX > len(matrix)-1 || newY < 0 || newY > len(matrix[0])-1 || (matrix[newX][newY] != char && matrix[newX][newY] != strings.ToLower(char)) {
			total++
		}
	}
	return total
}

func getUnvisitedNodes(matrix [][]string, coordinate Coordinate, char string) []Coordinate {
	coordinates := make([]Coordinate, 0)

	for _, direction := range directions {
		newX := coordinate.x + direction.x
		newY := coordinate.y + direction.y

		if newX >= 0 && newX < len(matrix[0]) && newY >= 0 && newY < len(matrix) && matrix[newX][newY] == char {
			coordinates = append(coordinates, Coordinate{newX, newY})
		}
	}

	return coordinates
}

func printMatrix(matrix [][]string) {
	fmt.Println()
	for _, row := range matrix {
		fmt.Println(row)
	}
	fmt.Println()
}

func partTwoSolution(matrix [][]string) int {
	total := 0

	for x := range matrix {
		for y := range matrix[x] {
			if !IsSingleLetterLowercase(matrix[x][y]) {
				area, perimeter := calculateRegionPartTwo(matrix, Coordinate{x, y})
				total += area * perimeter
			}
		}
	}

	return total
}




func sortCoordinates(coordinates []Coordinate) {
    sort.Slice(coordinates, func(i, j int) bool {
        if coordinates[i].x != coordinates[j].x {
            return coordinates[i].x < coordinates[j].x
        }
        return coordinates[i].y > coordinates[j].y
    })
}

func calculateRegionPartTwo(matrix [][]string, startingCoordinate Coordinate) (int, int) {
	area := 0

	unvisitedCoordinates := make([]Coordinate, 0)

	unvisitedCoordinates = append(unvisitedCoordinates, startingCoordinate)

	currentChar := matrix[startingCoordinate.x][startingCoordinate.y]

	outsideCoors := make([]Coordinate, 0)

	for len(unvisitedCoordinates) > 0 {
		for _, coor := range unvisitedCoordinates {
			outsideCoors = append(outsideCoors, getTouchingCoors(matrix, coor, currentChar)...)
			for _, c := range unvisitedCoordinates {
				matrix[c.x][c.y] = strings.ToLower(matrix[c.x][c.y])
			}
			area++
			unvisitedCoordinates = unvisitedCoordinates[1:]
			unvisitedCoordinates = append(unvisitedCoordinates, getUnvisitedNodes(matrix, coor, currentChar)...)
		}
	}

	perimeter := 0

	fmt.Println(currentChar, outsideCoors)

	sortCoordinates(outsideCoors)
	fmt.Println(outsideCoors)
	
	return area, perimeter
}



func getTouchingCoors(matrix [][]string, coor Coordinate, char string) []Coordinate {
	result := make([]Coordinate, 0)

	for _, direction := range directions {
		newX := coor.x + direction.x
		newY := coor.y + direction.y

		if newX < 0 || newX > len(matrix)-1 || newY < 0 || newY > len(matrix[0])-1 || (matrix[newX][newY] != char && matrix[newX][newY] != strings.ToLower(char)) {
			result = append(result, Coordinate{newX, newY})
		}
	}
	return result
}

