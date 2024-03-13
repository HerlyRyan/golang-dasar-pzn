package main

import ("fmt")

// Function return values
func getHello(name string) string{
	hello := "Hello " + name 
	return hello
}

// Function return multiple values
func getFullName() (string, string){
	return "Herly Riyanto", "Hidayat"
}

// Cara lain deklarasi function return multiple values
func getPerson()(fullName, address string, age int){
	fullName = "Herly Riyanto Hidayat"
	address = "Banjarbaru"
	age = 21
	return fullName, address, age
}

func main(){	
	result := getHello("Herly")
	fmt.Println(result)

	fmt.Println(getHello("Budi"))
	fmt.Println(getHello("Joko"))

	// Pemanggilan function getFullName()
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	namaAwal, _ := getFullName() // Cara mendapatkan value dari function getFullName() tanpa return value lastName
	fmt.Println(namaAwal)

	// 
	fullName, address, age := getPerson()
	fmt.Println("Hai namaku", fullName, "Aku tinggal di", address, "Umurku", age)
}