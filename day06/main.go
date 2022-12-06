package main

import "fmt"

func uniqueCounter(lines []string, uniqueCount uint8) int {
	input := lines[0]
	for i := 0; i < len(input)+1-int(uniqueCount); i += 1 {
		substring := (input[i : i+int(uniqueCount)])

		counter := make(map[string]int)
		for _, c := range substring {
			counter[string(c)]++
		}

		if len(counter) == int(uniqueCount) {
			return (i + int(uniqueCount))
		}

	}
	return 0
}

func main() {
	lines := getFileLines("input.txt")

	fmt.Printf("Solution to day 6 part one is %d \n", uniqueCounter(lines, 4))
	fmt.Printf("Solution to day 6 part two is %d", uniqueCounter(lines, 14))
}
