package main

import (
	"fmt"
)

func main(){
	var a = 20
	var b = 30
	var c = a + b
	fmt.Println(c)

	// Augmented assignments
	a += 5
	fmt.Println(a)
	a += 10
	fmt.Println(a)
}