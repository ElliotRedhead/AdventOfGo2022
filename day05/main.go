package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func crateMover(lines []string, reverseCrateDirection bool) {
	var readingCratePositions bool = true
	var crateStacks [9][]string
	var instructions [][]int
	for _, line := range lines {
		if readingCratePositions {
			for charIndex, char := range line {
				if char == rune('[') {
					crateStacks[((charIndex+4)/4)-1] = append(crateStacks[((charIndex+4)/4)-1], line[charIndex+1:charIndex+2])

				}
				if charIndex == 1 {
					if _, err := strconv.Atoi(string(char)); err == nil {
						readingCratePositions = false
					}
				}
			}
		} else if len(line) > 0 {
			re := regexp.MustCompile("[0-9]+")
			// [quantityToMove, initialPosition, finalPosition]
			var extractedNums = []int{}
			extractedNumStrings := re.FindAllString(line, -1)
			for _, numString := range extractedNumStrings {
				numInt, _ := strconv.Atoi(numString)
				extractedNums = append(extractedNums, numInt)
			}
			instructions = append(instructions, extractedNums)
		}
	}
	for _, instruction := range instructions {
		donorStack := make([]string, len(crateStacks[instruction[1]-1]))
		copy(donorStack, crateStacks[instruction[1]-1])
		recipientStack := crateStacks[instruction[2]-1]

		extracted := make([]string, len(donorStack[0:instruction[0]]))
		copy(extracted, donorStack[0:instruction[0]])
		if reverseCrateDirection {
			// reverse extracted
			for i, j := 0, len(extracted)-1; i < j; i, j = i+1, j-1 {
				extracted[i], extracted[j] = extracted[j], extracted[i]
			}
		}
		crateStacks[instruction[1]-1] = donorStack[instruction[0]:]
		crateStacks[instruction[2]-1] = append(extracted, recipientStack...)
	}

	var topCrates []string
	for _, stack := range crateStacks {
		if len(stack) > 0 {
			topCrates = append(topCrates, stack[0])
		}
	}

	fmt.Println(topCrates)

}

func main() {
	lines := getFileLines("input.txt")

	crateMover(lines, true)
	crateMover(lines, false)
}
