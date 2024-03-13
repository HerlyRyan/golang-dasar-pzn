package main

import (
	"fmt"
)

// Menggunakan variadic function dengan variadic argument parameter
func sumAll (numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// Menggunakan parameter biasa
func sumAllBiasa (numbers []int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main(){
	fmt.Println(sumAll(10, 10, 10)) // variadic
	fmt.Println(sumAllBiasa([]int{10,10, 10})) // biasa

	// Cara semisal kita punya data slice
	numbers := []int{0, 10, 90, 10}
	fmt.Println(sumAll(numbers...))
}