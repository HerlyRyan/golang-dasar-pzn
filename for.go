package main

import ("fmt")

func main() {
	// For init statement(i := 0);kondisi(i <= nilai);post statement(i++)
	var nilai int = 35
	for i := 0; i <= nilai; i++ {
		if i % 5 == 0 && i % 7 == 0 {
			fmt.Println("Fizz Buzz")
		} else if i % 5 == 0 {
			fmt.Println("Fizz")
		} else if i % 7 == 0{
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}

	// For biasa
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke",counter)
		counter++
	}

	// For Range
	names := []string{"Herly", "Riyanto", "Hidayat"}
	// manual
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	// Menggunakna for range
	for key, name := range names {
		fmt.Println("Index ke", key, "=", name)
	}
	// Menggunakan for each range
	for _, name := range names {
		fmt.Println("Halo ini tanpa key", name)
	}

	// Break dan Continue
	for i := 0; i < 10; i++ {
		fmt.Println("Perulangan ke", i)
		if i == 5 {
			break
		}
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		} else {
			fmt.Println("Perulangan ganjil ke", i)
		}
	}
}