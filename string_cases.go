package main

import (
	"fmt"
)

func capitalizer (s string) string {
	newStr := ""
	for _, x := range s {
		if x > 96 && x < 123 {
			newStr += string(x - 32)
		} else {
			newStr += string(x)
		}
	}
	return newStr
}

func lowercaser (s string) string {
	newStr := ""
	for _, x := range s {
		if x > 64 && x < 91 {
			newStr += string(x + 32)
		} else {
			newStr += string(x)
		}
	}
	return newStr
}

// doesn't handle special characters or numbers
func altercaser (s string) string {
	newStr := ""
	capital := true
	for _, x := range s {
		if (x > 96 && x < 123) || (x > 64 && x < 91) {
			capital = !capital
		} 
		if capital == true && (x > 96 && x < 123) {
			newStr += string(x - 32)
		} else if capital == false && (x > 64 && x < 91) {
			newStr += string(x + 32)
		} else {
			newStr += string(x)
		}
	}
	return newStr
}

func main() {
	fmt.Println(capitalizer("hello world")) // HELLO WORLD
	fmt.Println(lowercaser("HELLO HUMAN")) // hello human
	fmt.Println(altercaser("COMPUTERS are pretty NEAT, AREN'T they?")) // cOmPuTeRs ArE pReTtY nEaT, aReN't ThEy?
}
