package functions

func GetNextLetter(slice []string, index int) (rune, bool) {
	if index < 0 || index >= len(slice)-1 {
		return 0, false
	}
	for i := index + 1; i < len(slice); i++ {
		if slice[i] != "" {
			return rune(slice[i][0]), true
		}
	}
	return 0, false
}
