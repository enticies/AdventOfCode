package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	filePath := "example_input.txt"

	content, err := ioutil.ReadFile(filePath)

	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	inp := parseInput(content)

	// partOneResult := partOneSolution(inp)
	partTwoResult := partTwoSolution(inp)

	// fmt.Printf("Part one result is %d\n", partOneResult)
	fmt.Printf("Part two result is %d\n", partTwoResult)
}

func parseInput(content []byte) []int {
	s := make([]int, 0)

	splitContent := strings.Split(string(content), " ")

	for _, num := range splitContent {
		convNum, _ := strconv.Atoi(num)

		s = append(s, convNum)
	}

	return s
}

func copySlice(s []int) []int {
	newSlice := make([]int, 0)

	for _, val := range s {
		newSlice = append(newSlice, val)
	}

	return newSlice
}

func partOneSolution(nums []int) int {
	numberOfBlinks := 40

	newNums := copySlice(nums)

	for i := 0; i < numberOfBlinks; i++ {
		c := copySlice(newNums)
		newNums = newNums[:0]

		for _, num := range c {
			str := strconv.Itoa(num)

			if num == 0 {
				newNums = append(newNums, 1)
			} else if len(str)%2 == 0 {
				firstNum, _ := strconv.Atoi(str[:len(str)/2])
				secondNum, _ := strconv.Atoi(str[len(str)/2:])

				newNums = append(newNums, firstNum, secondNum)
			} else {
				newNums = append(newNums, num*2024)
			}
		}
	}

	return len(newNums)
}

func transform(num int) []int {
	str := strconv.Itoa(num)

	if num == 0 {
		return []int{1}
	} else if len(str)%2 == 0 {
		firstNum, _ := strconv.Atoi(str[:len(str)/2])
		secondNum, _ := strconv.Atoi(str[len(str)/2:])
		return []int{firstNum, secondNum}
	} else {
		return []int{num * 2024}
	}
}

func numberAfterBlinks(nums []int, numberOfBlinks int) []int {
	for blink := 0; blink < numberOfBlinks; blink++ {
		l := len(nums)
		for i := 0; i < l; i++ {
			num := nums[i]

			transformed := transform(num)

			if len(transformed) == 1 {
				nums[i] = transformed[0]
			} else {
				nums[i] = transformed[0]
				nums = append(nums, transformed[1])
			}

		}
		// fmt.Println(nums)
		// fmt.Println(nums, len(nums))
	}

	return nums
}

func partTwoSolution(nums []int) int {
	blinks := 10 

	precomputedOnes := make(map[int][]int, 0)

	precomputedOnes[0] = []int{1}

	for blink := 1; blink < blinks; blink++ {
        current := make([]int, len(precomputedOnes[blink-1]))
        copy(current, precomputedOnes[blink-1])
        
        result := numberAfterBlinks(current, 1)
        precomputedOnes[blink] = result 
    }

	fmt.Println(precomputedOnes)


	return len(numberAfterBlinks(nums, blinks))
}
