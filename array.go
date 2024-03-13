package main

import (
	"fmt"
)

func main(){
	var students [3]string // [3] adalah Batas kapasisas index array
	students[0] = "Herly"
	students[1] = "Riyanto"
	students[2] = "Hidayat"

	fmt.Println(students)
	fmt.Println(students[0])
	fmt.Println(students[1])
	fmt.Println(students[2])

	students[2] = "Hidayat 2.0" // Merubah isi array
	fmt.Println(students[2])

	var values = [4]int{90, 80, 70}
	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	var tes = [...]string{"Herly"}

	fmt.Println(tes)
	fmt.Println("Panjag array: ", len(tes)) // output panjang array
}