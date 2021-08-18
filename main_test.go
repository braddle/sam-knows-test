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
 	out, _ := cmd.CombinedOutput()

 	s.Contains(string(out), "requires input and output file locations")
}

func (s *EndToEndSuite) TestInputOne() {
	s.T().Skip("Handling some errors first")
	inputFile := "./inputs/1.json"
	outputFile := "./1.output"
	cmd := exec.Command("./sam", inputFile, outputFile)
 	err := cmd.Run()

 	s.NoError(err)
 	s.FileExists(outputFile)

 	expectedFile, _ := os.Open("./outputs/1.output")
 	actualFile, _ := os.Open(outputFile)

 	expected, _ := ioutil.ReadAll(expectedFile)
 	actual, _ := ioutil.ReadAll(actualFile)

 	s.Equal(string(expected), string(actual))
}

