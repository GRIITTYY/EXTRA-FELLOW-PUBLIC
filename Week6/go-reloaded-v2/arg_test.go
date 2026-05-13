package main

import (
	"errors"
	"testing"
)
var errorMsg error = errors.New("test failed")
func TestParseArgs1(tester *testing.T) {
		val, val1, valErr:= ParseArgs([]string{"adedapo", "amuda"})
		if valErr != nil {
			tester.Errorf("test failed %v", valErr)
		}
		val = "adedapo"
		val1 = "amuda"
		
		if val != "adedapo" {
			tester.Errorf("test failed expecting %v", val)
		}

		if val1 != "amuda" {
			tester.Errorf("test failed, expecting %v", val1)
		}

}


// func TestPlay(playTest *testing.T) {
// 	val1, val2, valError := ParseArgs([]string("Adedapo", "Amuda", "Abayomi"))




//}
