package basic

import "fmt"

type Cat struct {
}

func (d Cat) Sound() {
	fmt.Println("Meaw Meaw")
}

func (d Cat) NumOfLegs() {
	fmt.Println(4)
}
