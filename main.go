package main

import (
	"fmt"
	"os"
	"go-reloaded/functions" // Import your custom functions
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		fmt.Println("Usage: <inputFile> <outputFile>")
		return
	}

	inputFile := arguments[0]
	outputFile := arguments[1]

	// Read the input file
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error while reading the file:", err)
		return
	}

	// Process the content to handle (hex) and (bin)
	processedContent := handleCases(string(content))

	// Write the processed content to the output file
	err = os.WriteFile(outputFile, []byte(processedContent), 0644)
	if err != nil {
		fmt.Println("Error while writing to the file:", err)
	}
}

// handleCases processes the input content, replacing (hex) and (bin) values with their decimal equivalents
func handleCases(content string) string {
	runes := []rune(content) // Convert content to runes for safe character processing
	length := len(runes)
	result := make([]rune, 0, length)

	for i := 0; i < length; i++ {
		
		if i+4 < length && runes[i] == '(' && runes[i+1] == 'h' && runes[i+2] == 'e' && runes[i+3] == 'x' && runes[i+4] == ')' {
			// Found "(hex)", process the preceding word
			start := i - 1
			for start >= 0 && ((runes[start] >= '0' && runes[start] <= '9') || (runes[start] >= 'a' && runes[start] <= 'f') || (runes[start] >= 'A' && runes[start] <= 'F')) {
				start--
			}
			start++

			// Convert the hex string to decimal
			hexValue := string(runes[start:i])
			decimal := functions.HextoByte(hexValue)
			result = append(result, []rune(fmt.Sprintf("%d", decimal))...)
			i += 4 // Skip "(hex)"
		} else if i+4 < length && runes[i] == '(' && runes[i+1] == 'b' && runes[i+2] == 'i' && runes[i+3] == 'n' && runes[i+4] == ')' {
			// Found "(bin)", process the preceding word
			start := i - 1
			for start >= 0 && (runes[start] == '0' || runes[start] == '1') {
				start--
			}
			start++

			// Convert the binary string to decimal
			binValue := string(runes[start:i])
			decimal := functions.Bin(binValue)
			result = append(result, []rune(fmt.Sprintf("%d", decimal))...)
			i += 4 // Skip "(bin)"
		} else {
			// Append the current character to the result
			result = append(result, runes[i])
		}
	}

	return string(result)
}
