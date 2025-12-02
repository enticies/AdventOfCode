package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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
	// partTwoResult := partTwoSolution(copyMatrix(inp))

	fmt.Printf("Part one result is %d\n", partOneResult)
	// fmt.Printf("Part two result is %d\n", partTwoResult)
}

type Coordinate struct {
	x int
	y int
}

func parseInput(content []byte) [][]Coordinate {
	state := make([][]Coordinate, 0)

	splitContent := strings.Split(strings.TrimSpace(string(content)), "\n")

	for _, row := range splitContent {
		currentRobot := make([]Coordinate, 0)

		splitBySpace := strings.Split(strings.TrimSpace(row), " ")

		positionValues := strings.Split(strings.Split(splitBySpace[0], "=")[1], ",")
		velocityValues := strings.Split(strings.Split(splitBySpace[1], "=")[1], ",")

		positionX, _ := strconv.Atoi(positionValues[0])
		positionY, _ := strconv.Atoi(positionValues[1])

		velocityX, _ := strconv.Atoi(velocityValues[0])
		velocityY, _ := strconv.Atoi(velocityValues[1])

		currentRobot = append(currentRobot,
			Coordinate{positionX, positionY},
			Coordinate{velocityX, velocityY},
		)

		state = append(state, currentRobot)
	}

	return state
}

func moveRobot(rows int, cols int, positionX int, positionY int, velocityX int, velocityY int) (int, int) {
	positionX += velocityX
	positionY += velocityY

	if positionX < 0 {
		positionX = cols + positionX
	} else if positionX > cols-1 {
		positionX = positionX - cols
	}

	if positionY < 0 {
		positionY = rows + positionY
	} else if positionY > rows-1 {
		positionY = positionY - rows
	}

	return positionX, positionY
}

func printGrid(rows int, columns int, state [][]Coordinate) {
	grid := make([][]int, rows)

	for i := range grid {
		grid[i] = make([]int, columns)

		for j := range grid[i] {
			grid[i][j] = 0
		}
	}

	for _, robot := range state {
		pX, pY := robot[0].x, robot[0].y
		grid[pY][pX] += 1
	}

	for _, row := range grid {
		for _, col := range row {
			if col == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("@")
			}
		}
		fmt.Println()
	}

	fmt.Println()

}

func stateAfterOneSecond(currentState [][]Coordinate, rows int, cols int) [][]Coordinate {
	for i := 0; i < len(currentState); i++ {
		robot := currentState[i]

		positionX, positionY := robot[0].x, robot[0].y
		velocityX, velocityY := robot[1].x, robot[1].y

		newPositionX, newPositionY := moveRobot(rows, cols, positionX, positionY, velocityX, velocityY)

		robot[0].x = newPositionX
		robot[0].y = newPositionY

	}

	return currentState
}

func countRobots(rows int, cols int, state [][]Coordinate) int {
	firstQuadrantY, firstQuadrantX := rows/2, cols/2

	firstQuadrantTotal := 0
	secondQuadrantTotal := 0
	thirdQuadrantTotal := 0
	fourthQuadrantTotal := 0

	fmt.Println("FIRST: ", firstQuadrantY, firstQuadrantX)

	for _, robot := range state {
		x, y := robot[0].x, robot[0].y

		if x < firstQuadrantX && y < firstQuadrantY {
			firstQuadrantTotal += 1
		} else if x > firstQuadrantX && y < firstQuadrantY {
			secondQuadrantTotal += 1
		} else if y > firstQuadrantY && x < firstQuadrantX {
			thirdQuadrantTotal += 1
		} else if y > firstQuadrantY && x > firstQuadrantX {
			fourthQuadrantTotal += 1
		}
	}

	fmt.Println()
	fmt.Println("FIRST TOTAL: ", firstQuadrantTotal)
	fmt.Println("SECOND TOTAL: ", secondQuadrantTotal)
	fmt.Println("THIRD TOTAL: ", thirdQuadrantTotal)
	fmt.Println("FOURTH TOTAL: ", fourthQuadrantTotal)

	return firstQuadrantTotal * secondQuadrantTotal * thirdQuadrantTotal * fourthQuadrantTotal
}

func partOneSolution(state [][]Coordinate) int {
	seconds := 1000 
	rows := 7 
	cols := 11 

	for i := 0; i < seconds; i++ {
		fmt.Println(i)
		 printGrid(rows, cols, state)
		state = stateAfterOneSecond(state, rows, cols)
		time.Sleep(time.Millisecond * 200)
	}


	return countRobots(rows, cols, state)
}
