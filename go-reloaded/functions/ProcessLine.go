package functions

import (
	"fmt"
	"reflect"
	"strings"
)

func ProcessLine(textFile string) string {
	lines := strings.Split(string(textFile), "\n")
	var finalslice []string

	for _, line := range lines {

		worlds := strings.Fields(line)

		worlds = ProcessFlags(worlds)
		for {

			newWorlds := ProcessFlags(worlds)

			if reflect.DeepEqual(newWorlds, worlds) {
				worlds = ProcessFlags(worlds)
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
		filteredWords1 = ProcessFlags(filteredWords1)
		for {

			newWorlds := ProcessFlags(filteredWords1)

			if reflect.DeepEqual(newWorlds, filteredWords1) {
				filteredWords1 = ProcessFlags(filteredWords1)
				break
			}

			filteredWords1 = newWorlds
		}

		finalslice = append(finalslice, strings.Join(filteredWords1, " "))
		fmt.Println(finalslice)
	}
	slicepocess := []string{}
	finalfinalslice := []string{}
	for _, word := range finalslice {
		if word != "" {
			slicepocess = strings.Fields(word)
			finalfinalslice = append(finalfinalslice, strings.Join(slicepocess, " "))
		}
	}
	return strings.Join(finalfinalslice, "\n")
}
