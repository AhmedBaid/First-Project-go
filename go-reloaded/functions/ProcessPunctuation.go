package functions

import (
	"strings"
)

func ProcessPunctuation(str []string) []string {
	str = SplitByPunctuation(str)
	str = ProcessQuotes(str)

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
			if !inQuots && IsQuoteBalanced(str, i) {

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
	str = ConvertAToAn(str)
	return str
}
