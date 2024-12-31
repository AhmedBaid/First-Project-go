package functions

import (
	// "fmt"
	"strings"
)

func FuncPunc(line string) string {
	word := strings.Fields(line)
	// fmt.Println(word)
	for i := 0; i < len(word); i++ {
		count := 0

		switch {
		case word[i] == "." || word[i] == "," || word[i] == "!" || word[i] == "?" || word[i] == ":" || word[i] == ";":
			if i > 0 {
				for j := 0; j < i; j++ {
					if word[i-1-j] != "" {
						word[i-1-j] += word[i]
						word[i] = ""

					}
				}
			}

		case strings.HasPrefix(word[i], "."), strings.HasPrefix(word[i], ","), strings.HasPrefix(word[i], "!"), strings.HasPrefix(word[i], "?"), strings.HasPrefix(word[i], ":"), strings.HasPrefix(word[i], ";"):
			for j := 0; j < len(word[i]); j++ {
				if string(word[i][j]) == "." || string(word[i][j]) == "," || string(word[i][j]) == "!" || string(word[i][j]) == "?" || string(word[i][j]) == ":" || string(word[i][j]) == ";" {
					count++
				} else {
					break
				}
			}
			if i > 0 {
				for j := 0; j < i; j++ {
					if word[i-1-j] != "" {
						word[i-1-j] += word[i][:count]
						word[i] = word[i][count:]
					}
				}
			}

		}
	}
	filteredWords := []string{}
	for _, char := range word {
		if char != "" {
			filteredWords = append(filteredWords, char)
		}
	}

	return strings.Join(filteredWords, " ")
}
