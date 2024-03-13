package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Herly",
		"address": "Banjarbaru",
		"age": "21",		
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])	
	fmt.Println(person["age"])	

	delete(person, "address")
	fmt.Println(person)
}