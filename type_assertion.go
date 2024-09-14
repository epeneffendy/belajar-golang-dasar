package main

import (
	"fmt"
)

//sebuah cara untuk merubah tipe data tertentu menjadi tipe data yang diinginkan
//biasanya berguna saat berurusan dengan interfae kosong

func main() {
	// membuat variabel name dengan tipe interface kosong
	var name interface{}

	// mengisi variabel dengan data string,
	// saat ini variable masih bertipe interface kosong
	name = "Effendy"

	// melakukan type assertions,
	// merubah tipe interface kosong menjadi string
	nameString := name.(string)
	fmt.Println(nameString)

	//Switch Expression Untuk Type Assertions
	var age interface{}

	age = 20

	switch dataType := age.(type) {
	case string:
		fmt.Println(dataType, "adalah string")
	case int:
		fmt.Println(dataType, "adalah int")
	case bool:
		fmt.Println(dataType, "adalah bool")
	default:
		fmt.Println("belum tau")
	}

}
