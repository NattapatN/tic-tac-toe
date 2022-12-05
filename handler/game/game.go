package game

func CheckWinner(value string, gameBoard *[3][3]string) bool {
	convertBoard := convertToIntArray(value, gameBoard)
	var result bool
	for _, v := range winnerLogic {
		result = true
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if v[i][j] == 1 && convertBoard[i][j] != 1 {
					result = false
				}
			}
		}
		if result {
			return true
		}
	}

	return result
}

func CheckDue(gameBoard *[3][3]string) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if gameBoard[i][j] == " " {
				return false
			}
		}
	}
	return false
}
