package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"slices"
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

func partOneSolution(line string) int {
	diskMap := buildDiskMap(line)

	representation := buildDiskMapRepresentation(line, diskMap)

	movedFiles := moveFiles(representation)

	return checksum(movedFiles)
}

func buildDiskMapRepresentation(line string, diskMap map[int]Layout) []string {
	representation := make([]string, 0)

	for id, i := 0, 0; i < len(line); i += 2 {
		for blocks := 0; blocks < diskMap[id].blocks; blocks++ {
			representation = append(representation, fmt.Sprint(id))
		}

		for emptySpaces := 0; emptySpaces < diskMap[id].emptyBlocks; emptySpaces++ {
			representation = append(representation, ".")
		}

		id++
	}

	return representation
}

func parseInput(content []byte) string {
	return strings.TrimSpace(string(content))
}

type Layout struct {
	blocks      int
	emptyBlocks int
}

func moveFiles(representation []string) []string {
	emptySpaces := countEmptySpaces(representation)

	for i := 0; i < len(representation)-emptySpaces; i++ {
		if representation[i] == "." {
			for j := len(representation) - 1; j > 0; j-- {
				if representation[j] != "." {
					representation[i] = representation[j]
					representation[j] = "."
					break
				}
			}
		}
	}

	return representation
}

func findNextBlock(representation []string, start int) ([]string, int) {
	block := make([]string, 0)
	currentChar := representation[start]

	for ; start < len(representation); start++ {
		if currentChar == representation[start] {
			block = append(block, representation[start])
		} else {
			break
		}
	}

	return block, start - 1
}

func popBlock(r [][]string, start int) (int, []string) {
	for i := start; i > 0; i-- {
		if !slices.Contains(r[i], ".") {
			return i - 1, r[i]
		}
	}

	return -1, nil
}

func fillWithEmptySpace(block []string) {
	for i := range block {
		block[i] = "."
	}
}

func moveFilesPartTwo(representation []string) [][]string {
	copyOne := make([][]string, 0)
	copyTwo := make([][]string, 0)

	block, end := findNextBlock(representation, 0)

	copyOne = append(copyOne, block)
	copyTwo = append(copyTwo, block)


	for end < len(representation)-1 {
		block, end = findNextBlock(representation, end+1)
		copyOne = append(copyOne, block)
		copyTwo = append(copyTwo, block)
	}

	for right := len(copyOne) - 1; right > 0; right-- {
		for left := 0; left < right; left++ {
			if checkIfFits(copyOne[right], copyTwo[left]) {
				fitBlockIntoEmptySpace(copyOne[right], copyTwo[left])
				break
			}
		}
	}

	return copyTwo
}

func checkIfFits(from []string, into []string) bool {
	fromNumberOfEmptySpaces := 0

	for _, v := range from {
		if v == "." {
			fromNumberOfEmptySpaces++
		}
	}

	if fromNumberOfEmptySpaces == len(from) {
		return false
	}

	numberOfEmptySpaces := 0

	for _, v := range into {
		if v == "." {
			numberOfEmptySpaces++
		}
	}


	return len(from) <= numberOfEmptySpaces
}

func fitBlockIntoEmptySpace(from []string, into []string) []string {
	startingPosition := 0

	for i := 0; i < len(into); i++ {
		startingPosition = i

		if into[i] == "." {
			break
		}
	}

	for i := 0; i < len(from); i++ {
		into[startingPosition+i] = from[i]
	}

	fillWithEmptySpace(from)

	return into
}

func trimRepresentation(representation []string) []string {
	i := len(representation) - 1

	for ; representation[i] == "."; i-- {

	}

	return representation[:i+1]
}

func checksum(movedFiles []string) int {
	total := 0

	for i, val := range movedFiles {
		v, _ := strconv.Atoi(val)
		total += i * v
	}

	return total
}

func countEmptySpaces(line []string) int {
	total := 0

	for _, val := range line {
		if val == "." {
			total++
		}
	}

	return total
}

func printDiskMap(diskMap map[int]Layout) {
	fmt.Println("{")
	keys := make([]int, 0, len(diskMap))
	for k := range diskMap {
		keys = append(keys, k)
	}

	for i, key := range keys {
		layout := diskMap[key]
		fmt.Printf("  %d: {\n    \"blocks\": %d,\n    \"emptyBlocks\": %d\n  }",
			key, layout.blocks, layout.emptyBlocks)

		if i < len(keys)-1 {
			fmt.Println(",")
		} else {
			fmt.Println()
		}
	}
	fmt.Println("}")
}

func buildDiskMap(line string) map[int]Layout {
	diskMap := make(map[int]Layout)

	for id, i := 0, 0; i < len(line); i += 2 {
		num, _ := strconv.Atoi(string(line[i]))

		blocks := num
		emptyBlocks := -1

		if i+1 > len(line)-1 {
			emptyBlocks = 0
		} else {
			num, _ = strconv.Atoi(string(line[i+1]))
			emptyBlocks = num
		}

		currentFile := Layout{
			blocks:      blocks,
			emptyBlocks: emptyBlocks,
		}

		diskMap[id] = currentFile
		id++
	}

	return diskMap
}

func flatten(slice [][]string) []string {
	flattenedSlice := make([]string, 0)

	for _, s := range slice {
		for _, c := range s {
			flattenedSlice = append(flattenedSlice, string(c))
		}
	}

	return flattenedSlice
}

func partTwoSolution(line string) int {
	diskMap := buildDiskMap(line)

	representation := buildDiskMapRepresentation(line, diskMap)
	movedFiles := moveFilesPartTwo(representation)

	flattenedSlice := flatten(movedFiles)

	return checksum(flattenedSlice)
}
