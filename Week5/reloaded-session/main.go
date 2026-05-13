package main

import "fmt"

func main() {
	var userName string = "smomoh"
	var favouriteInt int = 10
	var phoneLength float64 = 6.43
	var isAfrican bool = false

	text := "My name is " + userName
	newNumber := 10 + 20 + favouriteInt
	newPhoneLength := 2 * phoneLength
	isNotAfrican := !isAfrican

	fmt.Println(text, "\n", newNumber, "\n", newPhoneLength, "\n", isNotAfrican)

	if phoneLength > 6.43 {
		fmt.Println("My phone is longer than the Apple Iphone 17 Pro Max.")
	} else if phoneLength < 6.43 {
		fmt.Println("The Apple Iphone 17 Pro Max is longer than my phone.")
	} else if phoneLength == 6.43 {
		fmt.Println("I think my phone is the Apple Iphone 17 Pro Max.")
	}

}
