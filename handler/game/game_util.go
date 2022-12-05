package game

var winnerLogic = [][][]int{
	{{1, 1, 1}, {0, 0, 0}, {0, 0, 0}},
	{{0, 0, 0}, {1, 1, 1}, {0, 0, 0}},
	{{0, 0, 0}, {0, 0, 0}, {1, 1, 1}},
	{{1, 0, 0}, {1, 0, 0}, {1, 0, 0}},
	{{0, 1, 0}, {0, 1, 0}, {0, 1, 0}},
	{{0, 0, 1}, {0, 0, 1}, {0, 0, 1}},
	{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
	{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}},
}

func convertToIntArray(value string, board *[3][3]string) [3][3]int {
	var result [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == value {
				result[i][j] = 1
			}
		}
	}
	return result
}
