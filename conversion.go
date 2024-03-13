package main

import (
	"fmt"
)

func main(){
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) // jadi minus karena melebihi batas dari int16

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	var nama = "Herly"
	var e uint8 = nama[0] // Membuat value dari variabel nama bisa dicari secara index
	var eString = string(e) // output H
	var tes = string(nama) // output Herly
	
	fmt.Println(e)
	fmt.Println(eString)
	fmt.Println(tes)
}