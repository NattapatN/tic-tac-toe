package player

func GetPlayerValue(player int) string {
	switch player {
	case 1:
		return "X"
	default:
		return "O"
	}
}
