package main

import "fmt"

// Challenge 1: First 3 Letters
// Task: Write a program that prints the first 3 letters of a given string.

// Example:

// Input: "Gopher"
// Output: "Gop"

// 🔧 Try it with both ASCII and Unicode characters!

func challenge1(str string) {
	first3Letters := str[:3]

	fmt.Printf("The first 3 letters of the word %q is %s\n", str, first3Letters)
}
