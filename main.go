package main

import (
	"encoding/json"
	"fmt"
	"github.com/braddle/sam-knows-test/data"
	"github.com/braddle/sam-knows-test/report"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Error: script requires input and output file locations to be provided")
		fmt.Println("\t ./sam <input file location> <output file location>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	in, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error: input file - No file found: %s", inputFile)
		os.Exit(2)
	}


	c := make([]data.Measure,0)
	b, _ := ioutil.ReadAll(in)
	err = json.Unmarshal(b, &c)
	if err != nil {
		fmt.Println("Error: Invalid file contents")
		os.Exit(3)
	}
	m := data.Measurements{c}
	s := report.Render(m)

	outputFile := os.Args[2]
	out, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
	if err != nil {
		fmt.Printf("Error: output file - File already exists: %s", outputFile)
		os.Exit(3)
	}

	out.Write([]byte(s))
	out.Close()
}
