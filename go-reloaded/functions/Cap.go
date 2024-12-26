package functions

func Capilize(s string) string {
	slice := []rune(s)
	inword := false
	for i, run := range slice {
		if (run >= 'a' && run <= 'z') || (run >= 'A' && run <= 'Z') || (run >= '0' && run <= '9') {
			if !inword {
				if run >= 'a' && run <= 'z' {
					slice[i] -= 32
				}
				inword = true
			} else {
				if run >= 'A' && run <= 'Z' {
					slice[i] += 32
				}
			}
		} else {
			inword = false
		}
	}
	return string(slice)
}
