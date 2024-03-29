package basic

import "fmt"

type Dog struct {
}

func (d Dog) Sound() {
	fmt.Println("Guk Guk")
}

func (d Dog) NumOfLegs() {
	fmt.Println(4)
}
