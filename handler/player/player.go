package player

import (
	"fmt"
	"strings"

	"github.com/NattapatN/tic-tac-toe/handler/board"
)

func PlayerInput(player int, gameBoard *[3][3]string) bool {
	var input string
	fmt.Printf("Player %v turn (%v): ", player, GetPlayerValue(player))
	fmt.Scan(&input)

	for !verifyPlayerInput(input) {
		if input == "exit" {
			return false
		}
		fmt.Println("Try again.")
		fmt.Printf("Player %v turn (%v): ", player, GetPlayerValue(player))
		fmt.Scan(&input)
	}

	//verify user input
	location := getLocation(input)
	value := GetPlayerValue(player)
	isFilled := board.FillBoard(value, location, gameBoard)
	if isFilled {
		return true
	} else {
		fmt.Println("Sorry you can't fil value in this field. Please try again.")
		return PlayerInput(player, gameBoard)
	}
}

func getLocation(input string) [2]int {
	splitInput := strings.Split(input, "")
	var col int
	var row int
	switch splitInput[0] {
	case "a":
		row = 0
	case "b":
		row = 1
	case "c":
		row = 2
	}
	switch splitInput[1] {
	case "1":
		col = 0
	case "2":
		col = 1
	case "3":
		col = 2
	}
	return [2]int{col, row}
}

func verifyPlayerInput(input string) bool {
	splitInput := strings.Split(input, "")
	if len(splitInput) != 2 {
		return false
	}
	switch splitInput[0] {
	case "a", "b", "c":
	default:
		return false
	}
	switch splitInput[1] {
	case "1", "2", "3":
	default:
		return false
	}
	return true
}
