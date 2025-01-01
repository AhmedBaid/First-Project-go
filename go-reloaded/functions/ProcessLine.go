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
		words := strings.Fields(line)
		words = ProcessFlags(words)
		for {
			newwords := ProcessFlags(words)
			if reflect.DeepEqual(newwords, words) {
				words = ProcessFlags(words)
				break
			}
			words = newwords
		}
		////////remove empty /////
		filteredWords := []string{}
		for _, word := range words {
			if word != "" {
				filteredWords = append(filteredWords, word)
			}
		}
		filteredWords = ProcessPunct(filteredWords)
		////////remove empty /////
		filteredWords1 := []string{}
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
