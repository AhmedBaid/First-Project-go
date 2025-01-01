package functions

import "unicode"

func Capilize(s string) string {
	slice := ""
	inword := true
	for _, run := range s {
		if unicode.IsLetter(run) && inword {
			slice += ToUpper(string(run))
			inword = false
		} else {
			slice += ToLower(string(run))
		}
	}
	return string(slice)
}
