package functions

import (
	"strings"
)

func checkNextLetter(slice []string, index int) (rune, bool) {
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

func checKer(text []string, index int) bool {
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

func fiitpunctuations(str []string) []string {
	result := []string{}
	currentWord := ""

	for _, word := range str {
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

func filltquot(str []string) []string {
	result := []string{}
	finresult := []string{}

	for _, word := range str {
		count := 0
		count1 := 0

		if len(word) == 1 {
			result = append(result, word)
			continue
		}
		if len(word) > 1 {
			if string(word[0]) == "'" && string(word[len(word)-1]) == "'" {
				if len(word) > 1 {
					for j := 0; j < len(word); j++ {
						if word[j] == '\'' {
							count++
							result = append(result, string(word[j]))
						} else {
							break
						}
					}
					for j := len(word) - 1; j >= 0; j-- {
						if word[j] == '\'' {
							count1++
						} else {
							break
						}
					}
					if count1 == len(word) {
						count1 = 0
					}
					if count != len(word) {

						result = append(result, word[count:len(word)-count1])
						for i := 0; i < count1; i++ {
							result = append(result, string(word[len(word)-1-i]))
						}
					}

				}
			} else if string(word[0]) == "'" {
				for j := 0; j < len(word); j++ {
					if word[j] == '\'' {
						count++
						result = append(result, string(word[j]))
					} else {
						break
					}
				}

				result = append(result, string(word[count:]))
			} else if string(word[len(word)-1]) == "'" {
				for j := len(word) - 1; j >= 0; j-- {
					if word[j] == '\'' {
						count++
					} else {
						break
					}
				}
				result = append(result, string(word[:len(word)-count]))
				for i := 0; i < count; i++ {
					result = append(result, string(word[len(word)-1]))
				}

			} else {
				result = append(result, string(word))
			}
		}
	}

	for i := 0; i < len(result); i++ {
		if result[i] == "'" {
			
				finresult = append(finresult, result[i])
			
		} else if result[i] != "'" {
			finresult = append(finresult, result[i])
		}
	}

	return finresult
}

func AtoAn(str []string) []string {
	for i := 0; i < len(str); i++ {
		switch {
		case len(str[i]) >= 1:

			letter, found := checkNextLetter(str, i)

			if str[i] == "'a" || str[i] == "a" || str[i] == "A" || str[i] == "'A" {

				for j := i; j < len(str); j++ {
					if letter != '\'' && str[j] != "" {
						letter, found = checkNextLetter(str, j)
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

func filter(str []string) []string {
	str = fiitpunctuations(str)
	str = filltquot(str)

	inQuots := false
	for i := 0; i < len(str); i++ {
		pass := true
		switch {
		case str[i] == "'":
			x := 0
			for k := i; k >= 0; k-- {
				if i > 0 {
					if str[k] == "" {
						x++
					} else {
						break
					}
				}
			}
			if !inQuots && checKer(str, i) {

				str[i] = "'" + str[i+1]
				if str[i+1] == "'" {
					pass = false
				} else {
					inQuots = true
				}
				str[i+1] = ""

			} else if inQuots && pass {
				for j := 0; j < i; j++ {
					if str[i-1-j] != "" {
						str[i-1-j] += "'"
						str[i] = ""
						inQuots = false
						break
					}
				}
			}
		case str[i] == "." || str[i] == "," || str[i] == "!" || str[i] == "?" || str[i] == ":" || str[i] == ";":
			if i > 0 {
				for j := 0; j < i; j++ {
					if str[i-1-j] != "" {
						str[i-1-j] += str[i]
						str[i] = ""

					}
				}
			}
		case strings.HasPrefix(str[i], ".") || strings.HasPrefix(str[i], ",") || strings.HasPrefix(str[i], "!") || strings.HasPrefix(str[i], "?") || strings.HasPrefix(str[i], ":") || strings.HasPrefix(str[i], ";"):
			var cot int
			passcot := true
			for j := 0; j < len(str[i]); j++ {
				if str[i][j] == '.' || str[i][j] == ',' || str[i][j] == '!' || str[i][j] == '?' || str[i][j] == ':' || str[i][j] == ';' {
					cot++
				} else {
					break
				}
			}
			if i > 0 && passcot {
				for j := i - 1; j >= 0; j-- {
					if str[j] != "" {

						str[j] += string(str[i][:cot])
						str[i] = str[i][cot:]
						passcot = false
						break
					}
				}
			}

		}
	}
	str = AtoAn(str)
	// fmt.Println("str = ", str)
	return str
}
