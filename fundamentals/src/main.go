package main

import (
	"fmt"
	pk "fundamentals/src/mypackage"
)

func main() {

	var myCar pk.CarPublic
	myCar.Brand = "Tesla"
	myCar.Year = 2002

	fmt.Println(myCar)

	pk.PrintMessage()
}
