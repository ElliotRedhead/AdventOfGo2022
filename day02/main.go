package main

import (
	"fmt"
	"strings"
)

// Part one: The key indicates OPOneNT_SELECTION PLAYER_SELECTION
// a b c - rock paper scissors (opOnent)
// x y z - rock paper scissors (player)

// Part two: The key indicates OPOneNT_SELECTION DESIRED_OUTCOME
// a b c - rock paper scissors (opOnent)
// x y z - lose draw win

func main() {

	lines := getFileLines("input.txt")

	choiceMap := map[string]uint8{"X": 1, "Y": 2, "Z": 3}
	winCombo := map[string]string{"X": "C", "Y": "A", "Z": "B", "C": "X", "A": "Y", "B": "Z"}
	drawCombo := map[string]string{"X": "A", "Y": "B", "Z": "C", "A": "X", "B": "Y", "C": "Z"}
	lossCombo := map[string]string{"X": "B", "Y": "C", "Z": "A", "B": "X", "C": "Y", "A": "Z"}

	var choiceScoreOne uint16
	var drawnScenariosOne uint16
	var wonScenariosOne uint16

	var choiceScoreTwo uint16
	var drawnScenariosTwo uint16
	var wonScenariosTwo uint16

	for _, line := range lines {
		scenario := strings.Split(line, " ")
		// [opponentSelection, playerSelection]
		choiceScoreOne += uint16(choiceMap[scenario[1]])
		if drawCombo[scenario[1]] == scenario[0] {
			drawnScenariosOne++
		} else if winCombo[scenario[1]] == scenario[0] {
			wonScenariosOne++
		}

		// [opponentSelection, desiredOutcome]
		if scenario[1] == "Z" {
			// win
			wonScenariosTwo++
			choiceScoreTwo += uint16(choiceMap[winCombo[scenario[0]]])
		} else if scenario[1] == "Y" {
			// draw
			drawnScenariosTwo++
			choiceScoreTwo += uint16(choiceMap[drawCombo[scenario[0]]])
		} else if scenario[1] == "X" {
			// loss
			choiceScoreTwo += uint16(choiceMap[lossCombo[scenario[0]]])
		}
	}
	fmt.Printf("drawn: %d, won: %d \n", drawnScenariosOne, wonScenariosOne)

	fmt.Printf("Solution to day two part one is %d. \n", choiceScoreOne+(drawnScenariosOne)*3+(wonScenariosOne)*6)
	fmt.Printf("Solution to day two part two is %d.", choiceScoreTwo+(drawnScenariosTwo)*3+(wonScenariosTwo)*6)
}
