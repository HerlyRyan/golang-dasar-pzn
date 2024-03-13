package main
import ("fmt")

func main() {
  var firstName string = "Herly"
  var lastName = "Riyanto Hidayat"
  var age int = 21
  isMarried := false

  fmt.Println("Nama saya adalah", firstName, lastName)
  fmt.Println("Umur saya adalah", age)

  if !isMarried {
    fmt.Println("Belum Menikah")
  } else {
    fmt.Println("Sudah menikah")
  }

  fmt.Println(len(firstName))
  fmt.Println(firstName[0]) // Output byte bukan karakter string
  fmt.Println("Herly"[0]) // Output byte bukan karakter string
}