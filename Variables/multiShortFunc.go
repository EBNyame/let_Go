package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multi Short Func
//
// 	1. Declare two variables using multiple short declaration syntax
//
//  2. Initialize the variables using `multi` function below
//
//  3. Discard the 1st variable's value in the declaration
//
//  4. Print only the 2nd variable
//
// NOTE
//  You should use `multi` function
//  while initializing the variables
//
// EXPECTED OUTPUT
//  4
// ---------------------------------------------------------

func multiShortFunc() {
	// ADD YOUR DECLARATIONS HERE
	//
	_, b := multi()

	// THEN UNCOMMENT THE CODE BELOW

	fmt.Println(b)
}

// multi is a function that returns multiple int values
func multi() (int, int) {
	return 5, 4
}
