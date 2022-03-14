package main

import "fmt"

func isPalindromo(text string) string {
	var textReverse string

	for i := len(text) - 1; i >= 0; i-- {
		textReverse += string(text[i])
	}

	if text == textReverse {
		return "Es palindromo"
	} else {
		return "No es un palindromo"
	}

}

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

	// Loop through arrays & slices
	for _, value := range slice {
		fmt.Println(value)
	}

	var palindromo string = "amor a roma"

	fmt.Println(palindromo, "es palindromo?", isPalindromo(palindromo))

}
