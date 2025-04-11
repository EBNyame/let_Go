package main

import (
	"fmt"
	"strings"
)

func main() {
	OneToTen()
	evenToTwenty()
	eachLetter()
	throughSlice()
}

func OneToTen() {
	// Print numbers from 1 to 10.
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func evenToTwenty() {
	// Print only even numbers between 1 and 20.
	fmt.Println(".....................................")
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}
}

func eachLetter() {
	//Loop through a string and print each letter.
	fmt.Println(".....................................")
	word := "Ghana"
	for _, ch := range word {
		fmt.Printf("Index: %c\n", ch)
	}
}

func throughSlice() {
	//Loop through a slice of fruits and print only those that start with "A".
	fmt.Println(".....................................")
	names := []string{"Kofi", "Ewura", "Ama", "Case", "Amina"}

	for _, name := range names {
		// fmt.Println(name)
		if strings.HasPrefix(name, "A") {
			fmt.Println(name)
		}
	}
}
