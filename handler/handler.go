package handler

import (
	"fmt"

	"github.com/NattapatN/tic-tac-teo/handler/board"
	"github.com/NattapatN/tic-tac-teo/handler/game"
	"github.com/NattapatN/tic-tac-teo/handler/player"
)

func Handler() {
	fmt.Println("Game starting...")
	gameBoard := board.Init()
	board.PrintBoard(gameBoard)
	gameIsEnd := false
	currentPlayer := 1
	for !gameIsEnd {
		var playerInput bool
		playerInput = player.PlayerInput(currentPlayer, &gameBoard)
		if !playerInput {
			gameIsEnd = true
		} else {
			board.PrintBoard(gameBoard)
			if game.CheckWinner(player.GetPlayerValue(currentPlayer), &gameBoard) {
				fmt.Printf("Player %v is Winner. 🏆🎉\n", currentPlayer)
				gameIsEnd = true
			} else if game.CheckDue(&gameBoard) {
				fmt.Printf("Game %v is Due. 🤝\n", currentPlayer)
				gameIsEnd = true
			}
			switch currentPlayer {
			case 1:
				currentPlayer = 2
			case 2:
				currentPlayer = 1
			}
		}
	}
	fmt.Println("Game is end.")
}
