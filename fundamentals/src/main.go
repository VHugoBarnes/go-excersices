package main

import "fmt"

func main() {

	//? Switch

	switch modulo := 4 % 2; modulo {
	case 0:
		fmt.Println("Even")
	default:
		fmt.Println("Odd")
	}

	//? Without condition
	value := 200
	switch {
	case value > 100:
		fmt.Println("Is gt than 100")
	case value < 0:
		fmt.Println("Is less than 0")
	default:
		fmt.Println("No condition")
	}

}
