package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgs(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * a
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {

	//? Functions
	normalFunction("Hello Keko")

	tripleArgs(149, 15, "Malia")

	var x int = returnValue(2)
	fmt.Println(x)

	var value1, value2 int = doubleReturn(2)
	fmt.Println(value1, value2)

	// If we just want one value
	var val1, _ int = doubleReturn(3)
	fmt.Println(val1)

}
