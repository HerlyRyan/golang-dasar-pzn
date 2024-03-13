package main

import "fmt"

type Filter func(string) string // Alias dari function filterName

func sayHelloName(name string, filter func(string) string) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func sayHelloNameWithAlias(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func filterName(name string)string{
	if name == "Anjing" {
		return "..."
	}
	return name
}

func main() {
	sayHelloName("Herly", filterName) // cara 1

	// cara 2
	filter := filterName
	sayHelloName("Anjing", filter)

	// Menggunakan alias
	sayHelloNameWithAlias("Herly", filterName) // cara 1

	// cara 2
	filterName := filterName
	sayHelloNameWithAlias("Anjing", filterName)
}