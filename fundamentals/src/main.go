package main

import "fmt"

func main() {

	//? Hash maps
	m := make(map[string]int)

	m["malia"] = 3
	m["keko"] = 3
	m["nicole"] = 22

	// loop map
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Find value
	// with ok we check if the value exists
	value, ok := m["malia"]

	fmt.Println(value, ok)

}
