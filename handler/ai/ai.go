package ai

import "fmt"

func AiFillTable(gameBoard [3][3]string) [2]int {
	fmt.Println("AI is thinking...")
	//init index
	indexBoard := [3][3]int{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gameBoard[i][j] == " " {
				indexBoard[i][j] = calculateIndex([2]int{i, j}, gameBoard)
				return [2]int{i, j}
			}
		}
	}
	return [2]int{}
}

func calculateIndex(location [2]int, gameBoard [3][3]string) int {

	return 0
}
