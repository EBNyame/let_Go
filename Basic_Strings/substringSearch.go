package main

import (
	"fmt"
	"strings"
)

func substringSearch() {
	sentence := "Go is the future of tech."
	text := "tech"

	found := strings.Contains(strings.ToLower(text), strings.ToLower(sentence))
	pos := strings.Index(sentence, text)
	cons := strings.Count(sentence, "o")
	fmt.Println(found)
	fmt.Println(pos)
	fmt.Println(cons)

}
