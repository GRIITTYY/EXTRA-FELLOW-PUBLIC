package main

import (
	"strconv"
	"strings"
)

func ParseHex(content string) string {
	splitContent := strings.Fields(content)
	
	for index, word := range splitContent {
		if word == "(hex)" {
			hexVal := splitContent[index-1]
			newValue, _ := strconv.ParseInt(hexVal, 16, 64)
			splitContent[index-1] = strconv.Itoa(int(newValue))
			splitContent = append(splitContent[:index], splitContent[index+1:]...)
		}
	}

	tranformedHex := strings.Join(splitContent, " ")
	return tranformedHex
}
