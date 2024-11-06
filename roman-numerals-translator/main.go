package main

import (
	"fmt"
	"strings"
)

func main() {
	romanNumerals := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	myRomanNum := ""
	fmt.Println("Enter a roman numeral")
	fmt.Scanf("%v", &myRomanNum)

	val := 0
	splitString := strings.Split(myRomanNum, "")
	for i := 0; i < len(splitString); i++ {
		_, exist := romanNumerals[splitString[i]]
		if exist {
			val += romanNumerals[splitString[i]]

		} else {
			fmt.Printf("Roman numeral does not exist")
		}
	}

	fmt.Printf("%d", val)

}
