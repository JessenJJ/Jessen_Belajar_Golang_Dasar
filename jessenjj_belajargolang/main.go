package main

import (
	"fmt"
	"jessenjj_belajargolang/basic"
)

func main() {
	basic.TryDefer()

	basic.Pointer()

	// var newpost basic.Post
	// newpost.Title = "judul artikel"
	// newpost.Content = "lorem ipsyunm"
	// newpost.Owner = "jessen jie"
	// fmt.Printf("%s\n %s\n %s\n", newpost.Title, newpost.Content, newpost.Owner)

	myuser := basic.NewUser("jessss", "jiiee", 256)

	// panggil constructor dari struct function NewPost
	mypost := basic.NewPost("Judul artikel", "kontent artikel", *myuser)

	// edit contentnya
	mypost.EditContent("loren impusns")
	fmt.Printf("%s\n %s\n %s\n", mypost.Title, mypost.Content, mypost.Owner.GetFullName())

	cat := basic.Cat{}
	dog := basic.Dog{}
	fish := basic.Fish{}
	fmt.Printf("%s %d", cat)
	fmt.Printf("%s %d", dog)
	fmt.Printf("%s %d", fish)

	basic.TryExit()
	// defer fmt.Println("Program selesai")
	// os.Exit(1)
	// manifest type
	// var/const namaVar tipeData
	var firstName string = "Jessen"
	// const pi float32 = 3.14
	firstName = "Jessen Jie"
	// pi = 22/7
	// type interface
	// namaVar := value
	age := 30
	isMarried := false
	fmt.Println(firstName, age, isMarried)
	fmt.Printf("Nama saya %s berumur %d dan status nikah saya %t\n", firstName, age, isMarried)

	// perulangan normal for
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("Looping ke", i)
	// }
	// var i int = 1
	// for i < 5 {
	// 	fmt.Println("looping ke", i)
	// 	i++
	// }

	// perulangan foreach
	drinks := []string{"soda", "juice", "tea", "coffee tuku"}
	for idx, val := range drinks[1:3] {
		fmt.Printf("Tersedia minuman %d\n %s\n", idx, val)
	}
	// for _, val := range drinks {
	// 	fmt.Printf("Tersedia minuman %s\n", val)
	// }

	// membuat array
	var buah [3]string // deklarasi type 1

	buah[0] = "mangga"
	buah[1] = "rambutan"
	buah[2] = "durian"

	// deklarasi type 2
	// var kendaraan = [3]string{"mobil", "motor", "sepeda"}

	// membuat slice
	// var buahan = []string{"duren", "jambu", "mangga", "jeruk"}

	// slice value
	// a := 2
	// b := a
	// a = 3
	// fmt.Println(a, b)

	// slice refrence itu pass by reference
	a := [3]int{1, 2, 3}
	b := a[1:]
	a[2] = 4
	fmt.Println(a, b)
	// jika array pass by value

	// var angka map[string]int

	// angka["satu"] = 1
	// angka["dua"] = 2

	// dataDiri := map[string]string{
	// 	"hooby": "main game",
	// 	"suka":  "main valorant",
	// }

	// Map
	// asosiasi, key-value

	// map butuh deklarasi saja, isinya nanti
	angka := make(map[string]int)
	angka["nama"] = 1
	angka["umur"] = 2

	// map perlu di isi langsung setelah deklarasi
	person := map[string]string{
		"nama": "jessen",
		"umur": "30",
	}

	fmt.Println(angka, person)

	//function
	hasil := penjumlahan(1, 2)
	fmt.Println(hasil, isEven(hasil))

	//unexported function call
	fmt.Println(isOdd(20), isOdd(3))

	fmt.Println(basic.GetFullName("a", "b"))

}

func penjumlahan(num1 int, num2 int) int {
	return num1 + num2
}

func isEven(num int) bool {
	if num == 0 {
		return false
	}
	if num%2 == 0 {
		return true
	}
	return false
}
