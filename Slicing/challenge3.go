package main

import "fmt"

// Task: Given an email address, extract and print the domain name (the part after @).

// Example:

// Input: "exodus@gmail.com"
// Output: "gmail.com"

func challenge3(email string) {
	maxLen := len(email)
	extraction := email[maxLen-9 : maxLen]
	fmt.Printf("hey!!, your email is %q but the domain name is %s\n", email, extraction)
}

// g m a i l . c o m
// 0 1 2 3 4 5 6 7 8
