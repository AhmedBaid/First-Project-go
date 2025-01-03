package functions

func SplitByPunctuation(slice []string) []string {
	result := []string{}
	currentWord := ""
	for _, word := range slice {
		for _, char := range word {
			if !(char == '.' || char == ',' || char == '!' || char == '?' || char == ':' || char == ';') {
				currentWord += string(char)
			} else {
				if currentWord != "" {
					result = append(result, currentWord)
					currentWord = ""
				}

				result = append(result, string(char))
			}
		}

		if currentWord != "" {
			result = append(result, currentWord)
			currentWord = ""
		}
	}
	return result
}
