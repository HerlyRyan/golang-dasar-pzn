package main

import ("fmt")

func getGoodBye(name string) string{
	return "Good Bye " + name
}

func main(){
	getgoodbye := getGoodBye
	fmt.Println(getgoodbye("Herly"))
}