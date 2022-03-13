package main

import "fmt"

func main() {

	//? Loops

	//? Conditional for
	for i := 0; i < 10; i++ {
		fmt.Println("i:", i)
	}

	fmt.Println()

	//? For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println()

	//? For forever
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 5 {
			break
		}
	}

}
