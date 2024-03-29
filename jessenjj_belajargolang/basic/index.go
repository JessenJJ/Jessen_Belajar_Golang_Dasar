package basic

import (
	"fmt"
	"os"
)

func TryDefer() {
	defer fmt.Println("Sampai jumpa")
	fmt.Println("Selamat datang")

}

func TryExit() {
	defer fmt.Println("Sampai jumpa")
	fmt.Println("Sampai jumpa")
	os.Exit(1)
}

func GetFullName(firstname string, Lastname string) string {
	return fmt.Sprintf("%s %s", firstname, Lastname)
}

func Pointer() {
	var tanggal int = 20
	var ptanggal *int = &tanggal
	fmt.Println(tanggal, *ptanggal)
	fmt.Println(&tanggal, ptanggal)
}

// belajar interface with cat and dogs

type IAnimal interface {
	Sound()
	NumOfLegs()
}

func Animal(animal IAnimal) {
	// cek suara bagaimana
	animal.Sound()
	// berapa jumlah kakinya
	animal.NumOfLegs()
}
