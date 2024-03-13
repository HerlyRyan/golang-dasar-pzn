package main

import ("fmt")

func sayHello(name string){
	fmt.Println("Hello", name)
}

func main(){
	data := []string{"Herly", "Yanto"}
	for i := 0; i < len(data); i++ {
		sayHello(data[i])
	}
}