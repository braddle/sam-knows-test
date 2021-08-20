package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: script requires input and output file locations to be provided")
		fmt.Println("\t ./sam <input file location> <output file location>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	_, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error: input file - No file found: %s", inputFile)
		os.Exit(2)
	}

	outputFile := os.Args[2]
	_, err = os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Printf("Error: output file - File already exists: %s", outputFile)
		os.Exit(3)
	}

}
