package functions

import (
	"reflect"
	"strings"
)

func ProcessLine(textFile string) string {
	lines := strings.Split(textFile, "\n")
	var FinalWordsSlice []string
	for _, line := range lines {
		words := strings.Fields(line)
		words = ProcessFlags(words)
		for {
			transformedWords := ProcessFlags(words)
			if reflect.DeepEqual(transformedWords, words) {
				break
			}
			words = transformedWords
		}
		// remove the empty string in the slice after processing of flags.
		filteredWords := []string{}
		for _, word := range words {
			if word != "" {
				filteredWords = append(filteredWords, word)
			}
		}
		filteredWords = ProcessPunctuation(filteredWords)
		// remove again the empty string in the slice after processing of punct.
		filterwords2 := []string{}
		for _, word := range filteredWords {
			if word != "" {
				filterwords2 = append(filterwords2, word)
			}
		}
		// apply the processflags until il nya pas un valid flag and remove empty
		filterwords2 = ProcessFlags(filterwords2)
		for {
			transformedWords := ProcessFlags(filterwords2)
			if reflect.DeepEqual(transformedWords, filterwords2) {
				break
			}
			filterwords2 = transformedWords
		}
		finalfilterwords := []string{}
		for _, word := range filterwords2 {
			if word != "" {
				finalfilterwords = append(finalfilterwords, word)
			}
		}
		// apply the processflags until il nya pas un valid flag  and remove empty
		finalfilterwords = ProcessFlags(finalfilterwords)
		for {
			transformedWords := ProcessFlags(finalfilterwords)
			if reflect.DeepEqual(transformedWords, finalfilterwords) {
				break
			}
			finalfilterwords = transformedWords
		}
		FinalWordsSlice = append(FinalWordsSlice, strings.Join(finalfilterwords, " "))
	}
	return strings.Join(FinalWordsSlice, "\n")
}
