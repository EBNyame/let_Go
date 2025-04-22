package main

import "fmt"



func main(){
	array()
	slice()
	slicingArray()
	addingToSlice()
}


// ............................Arrays................................
func array(){
	var x = [2]int{3, 6}
	x[0] = 5
	fmt.Println(x)
}


// ............................Slices................................
func slice(){
	var x = make([]int, 3)
	fmt.Println(x)
}

func slicingArray(){
	x := [5]int{3, 5, 6, 8, 9}
	y := x[1:5]
	fmt.Println(y)
}

func addingToSlice(){
	x := [] int{1, 3}
	x = append(x, 4,5,6,7)
	fmt.Println(x)

	for index, value := range x{
		fmt.Printf("Index:  %d, Value: %d\n", index, value)
	}
}