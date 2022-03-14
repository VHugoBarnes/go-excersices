package main

import "fmt"

func main() {

	//? Defer, break & continue

	// Defer
	// With defer the line will run last
	defer fmt.Println("Hello")
	fmt.Println("World")

	// Continue
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Is 2")
			continue
		}
	}

	for i := 0; i < 10; i++ {
		//continue
		if i == 2 {
			fmt.Println("llegue a 2")
			continue
		}
		//break
		if i == 8 {
			fmt.Println("llegue a 8")
			break
		}
		fmt.Println(i)
	}

	// Break

}
