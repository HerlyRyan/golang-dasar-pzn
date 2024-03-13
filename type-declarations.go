package main

import (
	"fmt"	
)

func main() {
	type NoKTP string // Membuat alias atau nama baru untuk type data string

	var namaKTP NoKTP = "Herly Riyanto Hidayat"
	var nomorKTP = 222222222222
	var convertNoKTP = NoKTP(nomorKTP) // Mengconvert no KTP ke string

	fmt.Println(namaKTP)
	fmt.Println(convertNoKTP)	
}