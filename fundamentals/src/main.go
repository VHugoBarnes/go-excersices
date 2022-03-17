package main

import "fmt"

type figure2D interface {
	area() float64
}

type square struct {
	base float64
}

type rectangle struct {
	base   float64
	height float64
}

func (sqr square) area() float64 {
	return sqr.base * sqr.base
}

func calculate(f figure2D) {
	fmt.Println("Area:", f.area())
}

func (rec rectangle) area() float64 {
	return rec.base * rec.height
}

func main() {
	mySquare := square{base: 2}
	myRectangle := rectangle{base: 4, height: 2}

	calculate(mySquare)
	calculate(myRectangle)

	// Lista interfaces
	myInterface := []interface{}{"Hola", 12, 102.3}
	fmt.Println(myInterface...)
}
