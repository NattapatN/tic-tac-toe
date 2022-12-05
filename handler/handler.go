package handler

import (
	"fmt"

	"github.com/fatih/color"
)

func Handler() {
	whilte := color.New(color.FgWhite)
	boldWhite := whilte.Add(color.Underline)
	boldWhite.Println(" |A B C")
	fmt.Print("1|")
	boldWhite.Println(" | | ")
	fmt.Print("2|")
	boldWhite.Println(" | | ")
	fmt.Print("3|")
	fmt.Println(" | | ")
}
