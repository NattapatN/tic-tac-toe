package main

import (
	"fmt"

	"github.com/NattapatN/tic-tac-toe/handler"
	"github.com/common-nighthawk/go-figure"
)

func main() {
	figure.NewFigure("Tic-Tac-Toe", "rectangles", true).Print()
	fmt.Println("Type 'exit' to exit game")
	fmt.Println("---------------------------------------------")
	handler.Handler()
}
