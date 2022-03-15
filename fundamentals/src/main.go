package main

import "fmt"

type pc struct {
	ram       int
	disk_size int
	brand     string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	a := 50
	b := &a

	*b = 100

	fmt.Println(b)
	fmt.Println(a)

	myPC := pc{ram: 16, disk_size: 1000, brand: "HP"}
	fmt.Println(myPC)

	myPC.ping()
	fmt.Println(myPC)

	myPC.duplicateRAM()
	fmt.Println(myPC)

	myPC.duplicateRAM()
	fmt.Println(myPC)

}
