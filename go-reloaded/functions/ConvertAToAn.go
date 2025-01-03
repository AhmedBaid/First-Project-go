package functions

import "strings"

func ConvertAToAn(str []string) []string {
	for i := 0; i < len(str); i++ {
		if str[i] == "a" || str[i] == "'a" || str[i] == "A" || str[i] == "'A" {
			letter, found := GetNextLetter(str, i)
			if found && strings.ContainsRune("aeiouhAEIOUH", letter) {
				str[i] += "n"
			}
		}
	}
	return str
}
