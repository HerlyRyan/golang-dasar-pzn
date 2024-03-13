package main

import "fmt"

func main() {
	const name = "Herly Riyanto Hidayat"
	const (
		firstName = "Herly Riyanto"
		lastName = " Hidayat"
	)
	fmt.Println(name)
	fmt.Println(firstName + lastName)
}