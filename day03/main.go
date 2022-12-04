package main

import "fmt"

func partOne(lines []string) {

	var matchingItems []rune
	for _, line := range lines {
		var rucksackCompartments [2]string
		rucksackCompartments[0] = line[0 : (len(line)+1)/2]
		rucksackCompartments[1] = line[(len(line)+1)/2:]
		matched := false
		for _, itemA := range rucksackCompartments[0] {
			if !matched {
				for _, itemB := range rucksackCompartments[1] {
					if itemA == itemB {
						matchingItems = append(matchingItems, itemA)
						matched = true
						break
					}
				}
			}
		}
	}

	var matchingPrioritySum uint32 = 0
	for _, matchingRune := range matchingItems {
		if matchingRune > 95 {
			// a -> z = 1 -> 26
			matchingPrioritySum += uint32(matchingRune) - 96
		} else {
			// A -> Z = 27 -> 52
			matchingPrioritySum += uint32(matchingRune) - 38
		}
	}
	fmt.Printf("Solution for day three part one is %d. \n", matchingPrioritySum)

}

func partTwo(lines []string) {
	var groupedRucksacks [][]string
	var linesSubset []string
	var matchingItems []rune

	for lineIndex, rucksack := range lines {
		// Remove duplicates
		allKeys := make(map[string]bool)
		list := []rune{}
		for _, item := range rucksack {
			if _, value := allKeys[string(item)]; !value {
				allKeys[string(item)] = true
				list = append(list, item)
			}
		}
		filteredRucksack := list
		linesSubset = append(linesSubset, string(filteredRucksack))

		// Group into threes
		if (lineIndex+1)%3 == 0 {
			groupedRucksacks = append(groupedRucksacks, linesSubset)
			linesSubset = nil
		}
	}

	for _, group := range groupedRucksacks {
		rucksackA := group[0]
		rucksackB := group[1]
		rucksackC := group[2]
		for _, itemA := range rucksackA {
			for _, itemB := range rucksackB {
				if itemA == itemB {
					for _, itemC := range rucksackC {
						if itemB == itemC {
							matchingItems = append(matchingItems, itemA)
							break
						}
					}
					break
				}
			}
		}
	}

	var matchingPrioritySum uint32 = 0
	for _, matchingRune := range matchingItems {
		if matchingRune > 95 {
			// a -> z = 1 -> 26
			matchingPrioritySum += uint32(matchingRune) - 96
		} else {
			// A -> Z = 27 -> 52
			matchingPrioritySum += uint32(matchingRune) - 38
		}
	}
	fmt.Printf("Solution for day three part two is %d. \n", matchingPrioritySum)

}

func main() {
	lines := getFileLines("input.txt")
	partOne(lines)
	partTwo(lines)
}
