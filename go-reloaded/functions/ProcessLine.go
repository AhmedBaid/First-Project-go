package functions

import (
	"strings"
)

func ProcessLine(content string) string {
	lines := strings.Split(content, "\n")
	var slice []string
	for _, line := range lines {
		Flags := ProcessFlags(line)
		punc:=FuncPunc(Flags)
		slice = append(slice,punc)
	}
	return strings.Join(slice, "\n")
}
