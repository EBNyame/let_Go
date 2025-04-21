package main

import (
	"fmt"
)

// func main() {

// 	// fmt.Printf("%#v\n", os.Args)
// 	msg := os.Args[1]
// 	l := len(msg)
// 	// ed := msg + strings.Repeat("!", l)
// 	rS := strings.Repeat("!", l)

// 	fmt.Println(rS + " " + msg + " " + rS)

// }

type Student struct {
	Name  string
	Level int
}

func main() {

	exodus := &Student{
		Name:  "Exodus",
		Level: 200,
	}

	level(exodus)

	fmt.Println(exodus)
}

func level(x *Student) {
	x.Level += 100
}
