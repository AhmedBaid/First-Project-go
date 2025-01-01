package functions

import "fmt"

func ConvertAToAn(str []string) []string {
	fmt.Println(str)
	for i := 0; i < len(str); i++ {
		switch {
		case len(str[i]) >= 1:
			letter, found := GetNextLetter(str, i)
			if str[i] == "'a" || str[i] == "a" || str[i] == "A" || str[i] == "'A" {
				for j := i; j < len(str); j++ {
					if letter != '\'' && str[j] != "" {
						letter, found = GetNextLetter(str, j)
						break
					}
				}
				if found && i+1 < len(str) && (letter == 'a' || letter == 'e' || letter == 'i' || letter == 'o' || letter == 'u' || letter == 'h') || (letter == 'A' || letter == 'E' || letter == 'I' || letter == 'O' || letter == 'U' || letter == 'H') {
					str[i] += "n"
				}
			}
		}
	}
	return str
}
