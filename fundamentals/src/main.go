package main

import (
	"fmt"
	"fundamentals/src/mypackage"
)

func main() {

	var myCar mypackage.CarPublic = mypackage.CarPublic{Brand: "Tesla", Year: 2021}

	fmt.Println(myCar)

}
