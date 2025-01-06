package functions

import "strings"

func ConvertAToAn(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == "a" || slice[i] == "'a" || slice[i] == "A" || slice[i] == "'A" {
			letter, found := GetNextLetter(slice, i)
			if found && strings.ContainsRune("aeiouhAEIOUH", letter) {
				slice[i] += "n"
			}
		}
	}
	return slice
}
