package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multi Assign #2
//
//  1. Assign the correct values to the variables
//     to match to the EXPECTED OUTPUT below
//
//  2. Print the variables
//
// HINT
//  Use multiple Println calls to print each sentence.
//
// EXPECTED OUTPUT
//  Air is good on Mars
//  It's true
//  It is 19.5 degrees
// ---------------------------------------------------------

func multipleAssign2() {
	// UNCOMMENT THE CODE BELOW:

	var (
		planet string
		isTrue bool
		temp   float64
	)

	// ADD YOUR CODE BELOW
	planet, isTrue, temp = "Earth", true, 19.5
	fmt.Println("Air is good on " + planet)
	fmt.Println("It's ", isTrue)
	fmt.Println("It is ", temp, "degress")
}
