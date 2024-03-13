package main

import ("fmt")

// looping 
func factorialLoop(value int)int  {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

// recursive
func faktorialRecursive(value int)int{
	if value == 1 {
		return 1
	} else {
		return value * faktorialRecursive(value-1)
	}
}

func main(){
	fmt.Println(faktorialRecursive(3))
	fmt.Println(factorialLoop(3))
}