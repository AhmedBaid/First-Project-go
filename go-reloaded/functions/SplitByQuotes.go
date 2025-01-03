package functions

func SplitByQuotes(slice []string) []string {
	result := []string{}
	for _, word := range slice {
		count := 0
		count1 := 0
		if len(word) == 1 {
			result = append(result, word)
			continue
		}
		if len(word) > 1 {
			if string(word[0]) == "'" && string(word[len(word)-1]) == "'" {
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
	return result
}
