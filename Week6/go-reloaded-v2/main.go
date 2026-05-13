package main

import (
	"errors"
	"fmt"
	"os"
)


var errorMessage error = errors.New("Error: Please provide exactly one input file and one output file")
var fileError error= errors.New("Error: cannot read input file ")

func main() {
	inputArgs := os.Args[1:]

	inputFileName, _, err:= ParseArgs(inputArgs)
	if err != nil {
		fmt.Println(err)
		return

	}


	inputFileContent, fileErr := ParseInputFileContent(inputFileName)

	if fileErr != nil {
		fmt.Println(fileError)
		return
	}

	fmt.Println(inputFileContent)

}
