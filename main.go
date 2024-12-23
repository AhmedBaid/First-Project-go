package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]
	if len(arguments) > 3 {
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	content, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("error while reading the file")
	}
	fmt.Printf("%s",content)

	//  output file 
	err = os.WriteFile(outputFile, content, 0644)
	if err != nil {
		fmt.Printf("Error while writing to the file: %v\n", err)
		return
	}
}