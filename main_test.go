package main_test

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/suite"
)

type EndToEndSuite struct {
	suite.Suite
}

func TestEndToEndSuite(t *testing.T) {
	suite.Run(t, new(EndToEndSuite))
}

func (s *EndToEndSuite) TestNoArguments() {
 	cmd := exec.Command("./sam")
 	out, err := cmd.CombinedOutput()

	s.Error(err)
 	s.Contains(string(out), "requires input and output file locations")
}

func (s *EndToEndSuite) TestInputFileDoesNotExist() {
	inputFile := "inputs/3.json"
	cmd := exec.Command("./sam", inputFile, "./reports/3.output")
	out, err := cmd.CombinedOutput()

	s.Error(err)
	s.Contains(string(out), "No file found: " + inputFile)
}

func (s *EndToEndSuite) TestOutputFileDoesExist() {
	outputFile := "./outputs/2.output"
	cmd := exec.Command("./sam", "inputs/2.json", outputFile)
	out, err := cmd.CombinedOutput()

	s.Error(err)
	s.Contains(string(out), "File already exists: " + outputFile)
}

func (s *EndToEndSuite) TestInputOne() {
	inputFile := "inputs/1.json"
	outputFile := "reports/1.output"
	os.Remove(outputFile)
	cmd := exec.Command("./sam", inputFile, outputFile)
	out, err := cmd.CombinedOutput()

 	s.NoError(err)
	s.Empty(out)
 	s.FileExists(outputFile)

 	expectedFile, _ := os.Open("./outputs/1.output")
 	actualFile, _ := os.Open(outputFile)

 	expected, _ := ioutil.ReadAll(expectedFile)
 	actual, _ := ioutil.ReadAll(actualFile)

 	s.Equal(string(expected), string(actual))
}

func (s *EndToEndSuite) TestInputTwo() {
	inputFile := "inputs/2.json"
	outputFile := "reports/2.output"
	os.Remove(outputFile)
	cmd := exec.Command("./sam", inputFile, outputFile)
	out, err := cmd.CombinedOutput()

	s.NoError(err)
	s.Empty(out)
	s.FileExists(outputFile)

	expectedFile, _ := os.Open("./outputs/2.output")
	actualFile, _ := os.Open(outputFile)

	expected, _ := ioutil.ReadAll(expectedFile)
	actual, _ := ioutil.ReadAll(actualFile)

	s.Equal(string(expected), string(actual))
}

func (s *EndToEndSuite) TestInvalidInputFile() {
	cmd := exec.Command("./sam", "inputs/bad.json", "reports/bad.out")
	out, err := cmd.CombinedOutput()

	s.Error(err)
	s.Contains(string(out), "Invalid file contents")
}

func (s *EndToEndSuite) TestEmptyInputFile() {
	cmd := exec.Command("./sam", "inputs/empty.json", "reports/empty.out")
	out, err := cmd.CombinedOutput()

	s.Error(err)
	s.Contains(string(out), "Could not find any data")
}
