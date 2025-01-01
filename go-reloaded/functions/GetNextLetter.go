package functions

// GetNextLetter returns the first letter of the next non-empty word in the slice, starting from the given index.
func GetNextLetter(slice []string, index int) (rune, bool) {
	if index < 0 || index >= len(slice)-1 {
		return 0, false
	}
	if index == 0 {
		if slice[index] != "" {
			return rune(slice[index][0]), true
		}
	} else {
		for i := index + 1; i < len(slice); i++ {
			if slice[i] != "" {
				return rune(slice[i][0]), true
			}
		}
	}
	return 0, false
}
