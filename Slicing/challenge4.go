package main

import "fmt"

// Task: Given the string "💡GoLang💥", slice out only "Go" and print it.

// Hint: Convert the string to []rune first!

func challenge4(str string) {
	runes := []rune(str)

	// maxLen := len(runes)

	extraction := string(runes[1:3])

	fmt.Printf("the first two letters of the word %q is %s\n", str, extraction)

}
