package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitDelimiter(r rune) bool {
	return r == ',' || r == '-'
}

func partOne(lines []string) {

	var completeOverlapCount uint16 = 0
	var groupedBounds [][]uint16
	for _, line := range lines {
		var boundsStr []string = strings.FieldsFunc(line, splitDelimiter)

		var bounds []uint16
		for _, boundStr := range boundsStr {
			boundInt, _ := strconv.ParseUint(boundStr, 10, 16)
			bounds = append(bounds, uint16(boundInt))
		}
		groupedBounds = append(groupedBounds, bounds)
	}

	for _, bounds := range groupedBounds {
		if bounds[0] <= bounds[2] && bounds[1] >= bounds[3] || bounds[2] <= bounds[0] && bounds[3] >= bounds[1] {
			completeOverlapCount++
		}
	}

	fmt.Printf("Solution to day four part one is %d. \n", completeOverlapCount)
}

func partTwo(lines []string) {

	var partialOverlapCount uint16 = 0
	var groupedBounds [][]uint16
	for _, line := range lines {
		var boundsStr []string = strings.FieldsFunc(line, splitDelimiter)

		var bounds []uint16
		for _, boundStr := range boundsStr {
			boundInt, _ := strconv.ParseUint(boundStr, 10, 16)
			bounds = append(bounds, uint16(boundInt))
		}
		groupedBounds = append(groupedBounds, bounds)
	}

	for _, bounds := range groupedBounds {
		lowerBA := bounds[0]
		upperBA := bounds[1]
		lowerBB := bounds[2]
		upperBB := bounds[3]
		if lowerBA <= lowerBB && upperBA >= lowerBB || lowerBB <= lowerBA && upperBB >= lowerBA || upperBA >= upperBB && lowerBA <= lowerBB || upperBB >= upperBA && lowerBB <= lowerBA {
			partialOverlapCount++
		}
	}

	fmt.Printf("Solution to day four part two is %d. \n", partialOverlapCount)
}

func main() {
	lines := getFileLines("input.txt")
	partOne(lines)
	partTwo(lines)
}
