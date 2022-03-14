package main

import "fmt"

func main() {

	//? Arrays
	var arr [4]int

	fmt.Println(arr, len(arr), cap(arr))

	//? Slices
	slice := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Println(slice, len(slice), cap(slice))

	//? Métodos en el slice
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])

	//? Append
	slice = append(slice, 7)
	fmt.Println(slice)

	//? Append new list
	newSlice := []int{8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)

}
