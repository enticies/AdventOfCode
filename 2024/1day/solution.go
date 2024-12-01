package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	firstColumn, secondColumn := parseInput(content)

	partOneResult := partOneSolution(firstColumn, secondColumn)
	partTwoResult := partTwoSolution(firstColumn, secondColumn)

	fmt.Printf("Part one result is %d\n", partOneResult)
	fmt.Printf("Part two result is %d\n", partTwoResult)
}

func parseInput(content []byte) ([]int, []int) {
	lines := strings.Split(strings.TrimSpace(string(content)), "\n")

	firstColumn := make([]int, 0)
	secondColumn := make([]int, 0)

	for _, line := range lines {
		splitLine := strings.Fields(line)

		firstNum, err := strconv.Atoi(splitLine[0])

		if err != nil {
			panic(err)
		}

		secondNum, err := strconv.Atoi(splitLine[1])

		if err != nil {
			panic(err)
		}

		firstColumn = append(firstColumn, firstNum)
		secondColumn = append(secondColumn, secondNum)
	}

	sort.Ints(firstColumn)
	sort.Ints(secondColumn)


	return firstColumn, secondColumn
}

func partOneSolution(firstColumn []int, secondColumn []int) int {
	totalDistance := 0 

	for i := 0; i < len(firstColumn); i++ {
		totalDistance += int(math.Abs(float64(firstColumn[i]) - float64(secondColumn[i])))
	}

	return totalDistance
}

func partTwoSolution(firstColumn []int, secondColumn []int) int {
	set := make(map[int]bool)	

	result := 0

	for _, valOne := range firstColumn {
		if !set[valOne] {
			count := 0
			for _, valTwo := range secondColumn {
				if valOne == valTwo {
					count++
				}
			}


			result += valOne * count
		} else {
			set[valOne] = true
		} 
	}

	return result
}