package functions

func IsQuoteBalanced(text []string, index int) bool {
	if text[index] == "'" {
		if index+1 < len(text) {
			for j := index + 1; j < len(text); j++ {
				if text[j] == "'" {
					return true
				}
			}
		}
	}
	return false
}
