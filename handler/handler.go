package handler

import (
	"fmt"

	"github.com/NattapatN/tic-tac-toe/handler/ai"
	"github.com/NattapatN/tic-tac-toe/handler/board"
	"github.com/NattapatN/tic-tac-toe/handler/game"
	"github.com/NattapatN/tic-tac-toe/handler/player"
)

func Handler() {
	gameBoard := board.Init()
	board.PrintBoard(gameBoard)
	gameIsEnd := false
	isExit := false
	currentPlayer := 1
	for !gameIsEnd && !isExit {
		playerInput := true
		if currentPlayer == 1 {
			playerInput = player.PlayerInput(currentPlayer, &gameBoard)
		} else {
			location := ai.AiFillTable(gameBoard)
			board.FillBoard(player.GetPlayerValue(currentPlayer), location, &gameBoard)
		}
		if !playerInput {
			isExit = true
		} else {
			board.PrintBoard(gameBoard)
			if game.CheckWinner(player.GetPlayerValue(currentPlayer), &gameBoard) {
				fmt.Printf("Player %v is Winner. 🏆🎉\n", currentPlayer)
				fmt.Println("---------------------------------------------")
				gameIsEnd = true
			} else if game.CheckDue(&gameBoard) {
				fmt.Println("Game is Due. 🤝")
				fmt.Println("---------------------------------------------")
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
	if isExit {
		fmt.Println("Game is end.")
	} else {
		var isWantToTryAgain string
		board.PrintPlayAgainCoversation()
		fmt.Scan(&isWantToTryAgain)
		if isWantToTryAgain == "y" {
			Handler()
		}
	}
}
