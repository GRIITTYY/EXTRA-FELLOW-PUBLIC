package main

import (
	"fmt"
	"os"
)

func main() {
	// Retrieve all arguments from the command line
	args := os.Args

	inputFileName, outputFileName, err := parseArgs(args)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(inputFileName, outputFileName)

}
