package main

import "fmt"

type pc struct {
	ram       int
	disk_size int
	brand     string
}

func (myPC pc) String() string {
	return fmt.Sprintf("RAM: %dGB, ROM: %dGB, brand: %s", myPC.ram, myPC.disk_size, myPC.brand)
}

func main() {

	var myPC pc = pc{ram: 16, brand: "HP", disk_size: 1000}
	fmt.Println(myPC)

}
