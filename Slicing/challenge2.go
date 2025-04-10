package main

import "fmt"

// Challenge 2: Last 4 Letters
// Task: Print the last 4 characters of any string.

// Example:

// Input: "Programming"
// Output: "ming"

// Hint: Use len(str) to calculate where to start slicing.
func challenge2(str string) {
	maxiLen := len(str)
	last4 := str[maxiLen-4 : maxiLen]

	fmt.Println(maxiLen)

	fmt.Printf("The last four letters of the word %q is %s\n", str, last4)

}

// g h a n a
// 0 1 2 3 4
