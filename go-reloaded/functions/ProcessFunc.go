package functions

import (
	"strconv"
	"strings"
)

func ProcessFunc(content string) string {
	chars := strings.Fields(content)

	for i := 0; i < len(chars); i++ {
		switch chars[i] {
		case "(cap)":
			if i > 0 {
				chars[i-1] = Capilize(chars[i-1])
			}
			chars[i] = ""
		case "(hex)":
			if i > 0 {
				chars[i-1] = strconv.Itoa(HextoByte(chars[i-1]))
			}
			chars[i] = ""
		case "(bin)":
			if i > 0 {
				chars[i-1] = strconv.Itoa(Bin(chars[i-1]))
			}
			chars[i] = ""
		case "(up)":
			if i > 0 {
				chars[i-1] = ToUpper(chars[i-1])
			}
			chars[i] = ""
		case "(low)":
			if i > 0 {
				chars[i-1] = ToLower(chars[i-1])
			}
			chars[i] = ""
		case "(low,":
			if i > 0 {
				chars[i-1] = ToLower(chars[i-1])
			}
			chars[i] = ""
		}
	}
	return strings.Join(chars, " ")
}
