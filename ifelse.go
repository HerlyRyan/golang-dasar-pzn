package main
import ("fmt")

func main() {
	var username string = "John"
	var password string = "1234"

	if username == "John" && password == "1234" {
		fmt.Println("Anda berhasil login")
	} else {
		fmt.Println("Username / password salah")
	}

	name := "Herly"
	length := len(name)
	if length > 5 {
		fmt.Println("Nama terlalu Panjang")
	} else {
		fmt.Println("Yak pas")
	}
}