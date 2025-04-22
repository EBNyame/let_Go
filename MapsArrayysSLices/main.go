package main

import "fmt"



func main(){
	array()
}


// ............................Arrays................................
func array(){
	var x = [2]int{3, 6}
	x[0] = 5
	fmt.Println(x)
}
