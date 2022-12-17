package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

func (myPc pc) String() string {
	return fmt.Sprintf("\nRAM: %d GB, Disk: %d GB, Brand: %s\n", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, brand: "DELL", disk: 300}

	fmt.Println("Pc details", myPc)
}
