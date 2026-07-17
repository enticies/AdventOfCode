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

func parseInput(content []byte) []string {
	return strings.Split(strings.TrimSpace(string(content)), "\n")
}

func partOneSolution(lines []string) int {
	prios := make([][]string, 0)
	orderings := make([][]string, 0)

	isPrio := true

	for _, line := range lines {
		if len(line) == 0 {
			isPrio = false
			continue
		}

		if (isPrio) {
			prios = append(prios, strings.Split(line, "|"))
		} else {
			orderings = append(orderings, strings.Split(line, ","))
		}
	}

	m := make(map[int][]int)

	for _, prio := range prios {
		convertedPrioSecond, _ := strconv.Atoi(prio[1])
		convertedPrioFirst, _ := strconv.Atoi(prio[0])
		if _, exists := m[convertedPrioSecond]; !exists {
			m[convertedPrioSecond] = make([]int, 0)
			m[convertedPrioSecond] = append(m[convertedPrioSecond], convertedPrioFirst)
		} else {
			m[convertedPrioSecond] = append(m[convertedPrioSecond], convertedPrioFirst)
		}
	}
	
	total := 0

	for _, ordering := range orderings {
		middle, _ := strconv.Atoi(ordering[len(ordering) / 2])
		total += middle
		for i := len(ordering) - 1; i > 0; i-- {
			c, _ := strconv.Atoi(ordering[i])
			if !contains(m[c], ordering[i-1]) {
				total -= middle
				break
			}
		}
	}


	return total 
}

func contains(slice []int, item string) bool {
    for _, value := range slice {
		item, _ := strconv.Atoi(item)
        if value == item {
            return true
        }
    }
    return false
}



func partTwoSolution(lines []string) int {
	prios := make([][]string, 0)
	orderings := make([][]string, 0)

	isPrio := true

	for _, line := range lines {
		if len(line) == 0 {
			isPrio = false
			continue
		}

		if (isPrio) {
			prios = append(prios, strings.Split(line, "|"))
		} else {
			orderings = append(orderings, strings.Split(line, ","))
		}
	}

	m := make(map[int][]int)

	for _, prio := range prios {
		convertedPrioSecond, _ := strconv.Atoi(prio[1])
		convertedPrioFirst, _ := strconv.Atoi(prio[0])
		if _, exists := m[convertedPrioSecond]; !exists {
			m[convertedPrioSecond] = make([]int, 0)
			m[convertedPrioSecond] = append(m[convertedPrioSecond], convertedPrioFirst)
		} else {
			m[convertedPrioSecond] = append(m[convertedPrioSecond], convertedPrioFirst)
		}
	}
	
	total := 0

	for _, ordering := range orderings {
		for i := len(ordering) - 1; i > 0; i-- {
			c, _ := strconv.Atoi(ordering[i])
			if !contains(m[c], ordering[i-1]) {
				for i, val := range ordering {
					left := 0
					right := 0

					tempSlice := make([]string, len(ordering))
					copy(tempSlice, ordering)

					newSlice := append(tempSlice[:i], tempSlice[i+1:]...)

					for _, v := range newSlice {
						for _, prio := range prios {
							if v == prio[0] && val == prio[1] {
								right++
							} else if v == prio[1] && val == prio[0] {
								left++
							}
						}
					}

					if left == right {
						cval, _ := strconv.Atoi(val)
						total += cval
					}

				}
				break
			}
		}
	}


	return total
}
