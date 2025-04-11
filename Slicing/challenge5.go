package main

import "fmt"

func challenge5(str string) {
	maxLen := len(str)
	extraction := str[1 : maxLen-1]

	fmt.Printf("The actual word is %q and after the extraction is %s\n", str, extraction)
}
