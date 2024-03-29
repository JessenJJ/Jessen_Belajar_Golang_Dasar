package basic

import "fmt"

type Fish struct {
}

func (d Fish) Sound() {
	fmt.Println("......")
}

func (d Fish) NumOfLegs() {
	fmt.Println(0)
}
