package functions

import (
	"reflect"
	"strings"
)

func ProcessLine(content string) string {
	lines := strings.Split(content, "\n")
	var slice []string
	for _, line := range lines {
		words := ProcessFunc(line)
		for {
			newwords:=ProcessFunc(words)
			if reflect.DeepEqual(newwords, words) {
				words = ProcessFunc(words)
				break
			}
			words = newwords
		}
		slice = append(slice,words)
	}
	return strings.Join(slice, "\n")
}
