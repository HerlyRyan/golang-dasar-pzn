package main

import "fmt"

func main() {
	name := "Herly"
	switch name {
	case "Herly":
		fmt.Print("Halo Herly")
	default:
		fmt.Println("Hi, Boleh kenalan?")		
	}
}