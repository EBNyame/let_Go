package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go inanc
//
//    inanc
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

func printYourName() {
	// get your name from the command-line
	myName := os.Args[1]
	// and print it
	fmt.Println(myName)
}
