package functions

import "strings"

func ConvertAToAn(str []string) []string {
	for i := 0; i < len(str); i++ {
		// Check if the current word is "a", "A", "'a", or "'A"
		if str[i] == "a" || str[i] == "'a" || str[i] == "A" || str[i] == "'A" {
			// Find the next letter of the next non-empty word
			letter, found := GetNextLetter(str, i)

			// Add "n" only if the next letter is a vowel or 'h'
			if found && strings.ContainsRune("aeiouhAEIOUH", letter) {
				str[i] += "n"
			}
		}
	}
	return str
}
