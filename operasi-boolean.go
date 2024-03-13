package main

import (
	"fmt"
)

func main() {
	// &&, ||, ! = untuk keterbalikan nilai
	var a = 10
	var b = 20	
	var c = a > b && b > a // Operasi ini bernilai true maka hasilnya true, dan jika satu saja false, maka hasilnya false
	var d = a < b || b < a // Kedua nilai false, hasilnya false ,dan salah satu true, hasilnya true
	var e = a == b	
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(!e)	

	var nilaiAkhir = 90
	var nilaiAbsen = 80

	var lulus = nilaiAkhir > 70 && nilaiAbsen > 70
	fmt.Println(lulus)
}