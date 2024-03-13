package main

import (
	"fmt"	
)

func main(){
	// iniSlice := []int{} untuk membuat slice
	// iniArray := [...]int{} untuk membuat array

	months:= [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	slices1 := months[4:8] // Mendapatkan nilai dari index ke-4 sampai 8
	fmt.Println(slices1)
	fmt.Println(slices1[2])
	fmt.Println(len(slices1))
	fmt.Println(cap(slices1)) // Mengecek kapasitas
	fmt.Println(append(slices1, "Tes")) // Menambah data pada array
	fmt.Println(months)

	slices2 := months[:3] // Mendapatkan nilai dari index awal sampai 3
	fmt.Println(slices2)

	slices3 := months[3:] // Mendapatkan nilai dari index ke-3 sampai terakhir
	fmt.Println(slices3)
	
	slices4 := months[:] // Mendapatkan semua nilai index
	fmt.Println(slices4)
	
	// Membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Herly"
	newSlice[1] = "Yanto"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Menambah 
	newSlice2 := append(newSlice, "Hidayat")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// Mengcopy slice
	fromSlice := months[:]
	toSlice := make([]string , len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice) // copy(destination Copy, from Data Copy)
	
	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}