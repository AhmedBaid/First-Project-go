package functions

func ProcessPunctuation(slice []string) []string {
	slice = SplitByPunctuation(slice)
	slice = SplitByQuotes(slice)
	inQuots := false
	for i := 0; i < len(slice); i++ {
		switch {
		case slice[i] == "'":
			x := 0
			for k := i; k >= 0; k-- {
				if i > 0 {
					if slice[k] == "" {
						x++
					} else {
						break
					}
				}
			}
			if !inQuots && IsQuoteBalanced(slice, i) {
				slice[i] = "'" + slice[i+1]
				if slice[i+1] == "'" {
					inQuots = false
				} else {
					inQuots = true
				}
				slice[i+1] = ""

			} else if inQuots {
				for j := 0; j < i; j++ {
					if slice[i-1-j] != "" {
						slice[i-1-j] += "'"
						slice[i] = ""
						inQuots = false
						break
					}
				}
			}
		case slice[i] == "." || slice[i] == "," || slice[i] == "!" || slice[i] == "?" || slice[i] == ":" || slice[i] == ";":
			if i > 0 {
				for j := 0; j < i; j++ {
					if slice[i-1-j] != "" {
						slice[i-1-j] += slice[i]
						slice[i] = ""
					}
				}
			}
		}
	}
	slice = ConvertAToAn(slice)
	return slice
}
