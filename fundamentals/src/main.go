package main

import "fmt"

func main() {

	// fmt package

	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage, worldMessage)

	// Printf
	nombre := "Keko kaka"
	edad := 3

	fmt.Printf("%s tiene %d años\n", nombre, edad)

	// Sprintf
	message := fmt.Sprintf("%s tiene %d años", nombre, edad)

	fmt.Println(message)

	// Data type
	fmt.Printf("helloMessage: %T", helloMessage)

}
