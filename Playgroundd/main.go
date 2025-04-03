package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// fmt.Printf("%#v\n", os.Args)
	msg := os.Args[1]
	l := len(msg)
	// ed := msg + strings.Repeat("!", l)
	rS := strings.Repeat("!", l)

	fmt.Println(rS + " " + msg + " " + rS)

}
