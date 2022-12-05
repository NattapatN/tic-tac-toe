package board

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintBoard(board [3][3]string) {
	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.Underline)
	boldWhite.Println(" |a b c")
	for i := 0; i < 3; i++ {
		fmt.Printf("%v|", i+1)
		for j := 0; j < 3; j++ {
			if i+1 < 3 {
				boldWhite.Printf("%v", board[i][j])
			} else {
				fmt.Printf("%v", board[i][j])
			}
			if j+1 < 3 {
				fmt.Print("|")
			}
		}
		fmt.Println()
	}
}

func PrintPlayAgainCoversation() {
	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.Underline)
	fmt.Print("Do you want to try again [")
	boldWhite.Print("y")
	fmt.Print("es/")
	boldWhite.Print("N")
	fmt.Print("o]: ")
}
