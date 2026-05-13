package main


import "os"



func ParseInputFileContent(fileName string) (string, error){
	
	data, err := os.ReadFile(fileName)

	if err != nil {
		return "", fileError
	}

	return string(data), nil

}