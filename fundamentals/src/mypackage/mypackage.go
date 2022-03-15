package mypackage

import "fmt"

// CarPublic Car with public access
type CarPublic struct {
	Brand string
	Year  int
}

// CarPrivate Car with private access
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage print a message
func PrintMessage() {
	fmt.Println("Hello")
}

// PrintMessage print a message
func printMessage() {
	fmt.Println("Hello world")
}
