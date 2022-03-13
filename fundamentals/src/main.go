package main

import "fmt"

func main() {

	// Variable types

	// Integers
	var in int = 19419491
	fmt.Println("int 64/32 bits", in)

	var in8 int8 = -128
	in8 = 127
	fmt.Println("int 8 bits", in8)

	var in16 int16 = 13033
	fmt.Println("int 16 bits", in16)

	var in32 int32 = 13033
	fmt.Println("int 32 bits", in32)

	var in64 int16 = 13033
	fmt.Println("int 64 bits", in64)

	// Positive integers
	var pin uint = 1
	fmt.Println("possitive int 64/32 bits", pin)

	var pin8 uint8 = 1
	fmt.Println("possitive int 8 bits", pin8)

	var pin16 uint16 = 1
	fmt.Println("possitive int 16 bits", pin16)

	var pin32 uint32 = 1
	fmt.Println("possitive int 32 bits", pin32)

	var pin64 uint64 = 1
	fmt.Println("possitive int 64 bits", pin64)

	// Decimals
	var flotante32 float32 = 1.41441
	fmt.Println("float 32 bits", flotante32)

	var flotante64 float64 = 1.41441
	fmt.Println("float 64 bits", flotante64)

	// String
	var str string = ""
	fmt.Println("str", str)

	// Booleans
	var boolean bool = false
	fmt.Println("Boolean", boolean)

	// Complex numbers
	var com64 complex64 = 14
	fmt.Println("Complex 64", com64)

	var com128 complex128 = 10 + 8i
	fmt.Println("Complex 128", com128)
}
