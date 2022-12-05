package board

func Init() [3][3]string {
	return [3][3]string{{" ", " ", " "}, {" ", " ", " "}, {" ", " ", " "}}
}

func FillBoard(playerValue string, location [2]int, board *[3][3]string) bool {
	if verifyValueCanFillBoard(playerValue, location, board) {
		board[location[0]][location[1]] = playerValue
		return true
	} else {
		return false
	}
}

func verifyValueCanFillBoard(playerValue string, location [2]int, board *[3][3]string) bool {
	switch board[location[0]][location[1]] {
	case " ":
		return true
	default:
		return false
	}
}
