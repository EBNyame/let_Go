package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Windows Path
//
//  1. Change the following program
//  2. It should use a raw string literal instead
//
// HINT
//  Run this program first to see its output.
//  Then you can easily understand what it does.
//
// EXPECTED OUTPUT
//  Your solution should output the same as this program.
//  Only that it should use a raw string literal instead.
// ---------------------------------------------------------

func windowsPath() {
	// HINTS:
	// \\ equals to backslash character
	// \n equals to newline character

	path := `c:\program files\duper super\fun.txt\n +
		c:\program files\really\funny.png`
	fmt.Println(path)
}
