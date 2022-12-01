package main

import (
	"fmt"
	"sort"
	"strconv"
)

func groupCals(fileLines []string) []uint32 {
	var groupedCals []uint32
	var elfCalsSum uint32 = 0
	for lineIndex, calsStr := range fileLines {
		calsInt, convErr := strconv.ParseUint(calsStr, 10, 32)
		if convErr != nil {
			groupedCals = append(groupedCals, elfCalsSum)
			elfCalsSum = 0
			continue
		}
		elfCalsSum += uint32(calsInt)
		if lineIndex == len(fileLines)-1 {
			groupedCals = append(groupedCals, elfCalsSum)
		}
	}
	return groupedCals
}

func aggregateGreatestCals(elvesCals []uint32, sliceLen uint8) uint32 {
	var calAggregate uint32 = 0
	for index, cals := range elvesCals {
		if uint8(index) > sliceLen-1 {
			break
		}
		calAggregate += cals
	}
	return calAggregate
}

func main() {
	lines := getFileLines("input.txt")
	elvesCals := groupCals(lines)
	sort.Slice(elvesCals, func(i, j int) bool { return elvesCals[i] > elvesCals[j] })
	fmt.Printf("Answer to part one of day one is %d. \n", elvesCals[0])

	chonkyElvesCals := aggregateGreatestCals(elvesCals, 3)

	fmt.Printf("Answer to part two of day one is %d.", chonkyElvesCals)
}
