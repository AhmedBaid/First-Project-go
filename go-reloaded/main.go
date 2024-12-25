package main

import (
	"fmt"
	"os"
	"strings"
	"go-reloaded/functions"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) != 2 {
		fmt.Println("An error occurred while inserting arguments")
		return
	}

	inputFile := arguments[0]
	outputFile := arguments[1]
	if inputFile==outputFile{
		fmt.Println("the files should not be the same")
	}
	if !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("the last arguments should end with .txt")
		return
	}

	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error while reading the file:", err)
		return
	}
	finalContent := functions.ProcessFunc(string(content))

	err = os.WriteFile(outputFile, []byte(finalContent), 0o644)
	if err != nil {
		fmt.Println("Error while writing to the file:", err)
	}
}
