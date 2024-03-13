package main

import ("fmt")

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	// Cara 1
	blacklist := func(name string) bool {
		return name == "Anjing"
	}
	registerUser("Herly", blacklist)

	// Cara 2
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}