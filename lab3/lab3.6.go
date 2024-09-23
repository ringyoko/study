package main

import (
	"fmt"
)

func main() {

	strings := []string{
		"Scarecrow",
		"Sgt. Leland Coyle, OkCPD",
		"Marilyn Manson, male beauty standard.",
		"Closer",
		"doc",
	}

	var longestString string
	maxLength := 0

	for _, str := range strings {

		if len(str) > maxLength {
			maxLength = len(str)
			longestString = str
		}
	}

	fmt.Println("Longest string:", longestString)
	fmt.Println("Longest string's length:", maxLength)
}
