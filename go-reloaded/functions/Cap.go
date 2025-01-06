package functions

import "unicode"

func Capilize(s string) string {
	str := ""
	inword := true
	for _, run := range s {
		if unicode.IsLetter(run) && inword {
			str += ToUpper(string(run))
			inword = false
		} else {
			str += ToLower(string(run))
		}
	}
	return string(str)
}
