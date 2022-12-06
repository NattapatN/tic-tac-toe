package ai

import (
	"fmt"
	"math/rand"
)

func AiFillTable(gameBoard [3][3]string) [2]int {
	fmt.Println("AI is thinking...")
	//init available fields
	availableFields := [][2]int{}
	//find available fieds
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gameBoard[i][j] == " " {
				availableFields = append(availableFields, [2]int{i, j})
			}
		}
	}
	randomInput := rand.Intn(len(availableFields))
	return availableFields[randomInput]
}
