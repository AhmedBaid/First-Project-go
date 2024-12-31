package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

func hex(hex string) string {
	decimal, err := strconv.ParseInt(hex, 16, 0)
	if err != nil {
		return hex
	}
	return fmt.Sprintf("%d", decimal)
}

func bin(bin string) string {
	decimal, err := strconv.ParseInt(bin, 2, 0)
	if err != nil {
		return bin
	}
	return fmt.Sprintf("%d", decimal)
}

func up(world string) string {
	return strings.ToUpper(world)
}

func low(world string) string {
	return strings.ToLower(world)
}

func cap(world string) string {
	x := ""
	pass := true
	for _, chr := range world {
		if unicode.IsLetter(chr) {
			if pass {
				x += up(string(chr))
				pass = false
			} else {
				x += low(string(chr))
			}
		} else {
			x += string(chr)
		}
	}
	return x
}

func correctNumber(numb string) (int, error) {
	x := ""
	for r, i := range numb {
		if r != len(numb)-1 && string(i) != "" {
			x += string(i)
		}
	}

	return strconv.Atoi(x)
}

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
	x := 0
	for i := 0; i < len(result); i++ {
		if result[i] == "'" {
			if checKer(result, i) {
				x++
				finresult = append(finresult, result[i])
			} else if x%2 != 0 {
				finresult = append(finresult, result[i])
			}
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
	return str
}

func keys(slices []string) []string {
	worlds := slices
	for i := 0; i < len(worlds); i++ {
		switch {
		case worlds[i] == "(hex)":
			for j := 0; j < i; j++ {
				if i > 0 && worlds[i-1-j] != "" {
					worlds[i-1-j] = hex(worlds[i-1-j])
					break
				}
			}

			worlds[i] = ""
		case worlds[i] == "(bin)":
			for j := 0; j < i; j++ {
				if i > 0 && worlds[i-1-j] != "" {
					worlds[i-1-j] = bin(worlds[i-1-j])
					break
				}
			}
			worlds[i] = ""
		case worlds[i] == "(up)":
			if i > 0 {
				worlds[i] = ""
				passup := true
				for j := 0; j < i; j++ {
					if worlds[i-1-j] != "" && passup {
						worlds[i-1-j] = up(worlds[i-1-j])

						passup = false
					}
				}
			} else {
				worlds[i] = ""
			}

		case worlds[i] == "(low)":
			if i > 0 {
				worlds[i] = ""
				passlow := true
				for j := 0; j < i; j++ {
					if worlds[i-1-j] != "" && passlow {
						worlds[i-1-j] = low(worlds[i-1-j])

						passlow = false
					}
				}
			} else {
				worlds[i] = ""
			}
		case worlds[i] == "(cap)":
			if i > 0 {
				worlds[i] = ""
				passcap := true
				for j := 0; j < i; j++ {
					if worlds[i-1-j] != "" && passcap {
						worlds[i-1-j] = cap(worlds[i-1-j])

						passcap = false
					}
				}
			} else {
				worlds[i] = ""
			}
		case worlds[i] == "(up,":

			if i+1 < len(worlds) && worlds[i+1][len(worlds[i+1])-1] == ')' {
				countStr := worlds[i+1]
				count, err := correctNumber(countStr)
				x := 0
				if err != nil {
					fmt.Println("Error: invalid count :", worlds[i+1])
				} else {
					for k := i - 1; k >= 0; k-- {
						if worlds[k] == "" {
							x++
						} else {
							break
						}
					}
					if count+x >= i {
						count = i - x
					}

					for j := 0; j < count+x; j++ {
						if worlds[i-1-j] != "" && countStr[len(countStr)-1] == ')' {
							worlds[i-1-j] = up(worlds[i-1-j])
						}
					}

					worlds[i] = ""
					worlds[i+1] = ""

				}

			}

		case worlds[i] == "(low,":
			if i+1 < len(worlds) && worlds[i+1][len(worlds[i+1])-1] == ')' {
				countStr := worlds[i+1]
				count, err := correctNumber(countStr)
				var x int
				if err != nil {
					fmt.Println("Error: invalid count :", worlds[i+1])
				} else {
					for k := i - 1; k >= 0; k-- {
						if worlds[k] == "" {
							x++
						} else {
							break
						}
					}
					if count+x >= i {
						count = i - x
					}

					for j := 0; j < count+x; j++ {
						if worlds[i-1-j] != "" {
							worlds[i-1-j] = low(worlds[i-1-j])
						}
					}

					worlds[i] = ""
					worlds[i+1] = ""

				}

			}
		case worlds[i] == "(cap,":

			if i+1 < len(worlds) && len(worlds[i+1]) > 1 {
				if worlds[i+1][len(worlds[i+1])-1] == ')' {
					countStr := worlds[i+1]
					count, err := correctNumber(countStr)
					x := 0
					if err != nil {
						fmt.Println("Error: invalid count :", worlds[i+1])
					} else {
						for k := i - 1; k >= 0; k-- {
							if worlds[k] == "" {
								x++
							} else {
								break
							}
						}
						if count+x >= i {
							count = i - x
						}

						for j := 0; j < count+x; j++ {
							if worlds[i-1-j] != "" {
								worlds[i-1-j] = cap(worlds[i-1-j])
							}
						}
						worlds[i] = ""
						worlds[i+1] = ""

					}
				}
			}

		}
	}
	return worlds
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_filename> <output_filename>")
		return
	}
	inputfile := os.Args[1]
	outputfile := os.Args[2]

	if !strings.HasSuffix(outputfile, ".txt") {
		fmt.Println("Error: Output file must have a .txt extension 😂")
		return
	}

	textFile, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error reading input file 😒:", err)
		return
	}

	lines := strings.Split(string(textFile), "\n")

	var processedLines []string

	for _, line := range lines {

		worlds := strings.Fields(line)

		worlds = keys(worlds)
		for {

			newWorlds := keys(worlds)

			if reflect.DeepEqual(newWorlds, worlds) {
				worlds = keys(worlds)
				break
			}

			worlds = newWorlds
		}

		filteredWords := []string{}
		filteredWords1 := []string{}
		for _, word := range worlds {
			if word != "" {
				filteredWords = append(filteredWords, word)
			}
		}

		filteredWords = filter(filteredWords)

		for _, word := range filteredWords {
			if word != "" {
				filteredWords1 = append(filteredWords1, word)
			}
		}
		filteredWords1 = keys(filteredWords1)
		for {

			newWorlds := keys(filteredWords1)

			if reflect.DeepEqual(newWorlds, filteredWords1) {
				filteredWords1 = keys(filteredWords1)
				break
			}

			filteredWords1 = newWorlds
		}

		processedLines = append(processedLines, strings.Join(filteredWords1, " "))

	}

	err = os.WriteFile(outputfile, []byte(strings.Join(processedLines, "\n")), 0o644)
	if err != nil {
		fmt.Printf("Error writing to output file 😱: %v\n", err)
		return
	}

	fmt.Println("Processing completed successfully!")
}
