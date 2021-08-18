package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Error: script requires input and output file locations to be provided")
	fmt.Println("\t ./sam <input file location> <output file location>")
	os.Exit(1)
}
