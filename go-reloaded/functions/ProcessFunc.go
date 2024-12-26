package functions

import (
	"strconv"
	"strings"
)

func ProcessFunc(content string) string {
	chars := strings.Fields(content)
	for i := 0; i < len(chars); i++ {
		switch {
		case chars[i] == "(hex)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = strconv.Itoa(HextoByte(chars[i-j-1]))
						chars[i] = ""
					}
				}
			} else {
				chars[i] = ""
			}
		case chars[i] == "(bin)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = strconv.Itoa(Bin(chars[i-j-1]))
						chars[i] = ""
					}
				}
			} else {
				chars[i] = ""
			}
		case chars[i] == "(cap)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = Capilize(chars[i-j-1])
						chars[i] = ""
					}
				}
			} else {
				chars[i] = ""
			}
		case chars[i] == "(up)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = ToUpper(chars[i-j-1])
						chars[i] = ""
					}
				}
			} else {
				chars[i] = ""
			}
		case chars[i] == "(low)":
			pass := true
			if i > 0 {
				for j := 0; j < i; j++ {
					if chars[i-j-1] != "" && pass {
						chars[i-j-1] = ToLower(chars[i-j-1])
						chars[i] = ""
					}
				}
			} else {
				chars[i] = ""
			}
		case strings.Contains(chars[i], "(low,") && strings.HasSuffix(chars[i], ")"):
			parts := strings.Split(chars[i], ",")
			if len(parts) == 2 && strings.HasPrefix(parts[0], "(low") {
				num, err := strconv.Atoi(strings.TrimSuffix(parts[1], ")"))
				if err == nil && i > 0 {
					count := i
					if num > count {
						for j := 0; j < count; j++ {
							chars[j] = ToLower(chars[j])
						}
					} else {
						for j := count - 1; j >= count-num; j-- {
							chars[j] = ToLower(chars[j])
						}
					}
					chars[i] = ""
				}
			}
		}
	}
	return strings.Join(chars, " ")
}
