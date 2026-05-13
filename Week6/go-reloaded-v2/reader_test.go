package main

import (
	"testing"
)

func TestParseInputfileContent(tester *testing.T) {
	fileContent, err := ParseInputFileContent("test.txt")
	if err != nil {
		tester.Fail()
		tester.Log("invalid content")
	}

	if fileContent != "Dapo" {

		tester.Fail()

		tester.Log("Invalid input content! expected", fileContent)
		
	}
}
