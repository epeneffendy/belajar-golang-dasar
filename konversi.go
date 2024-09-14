package main

import "fmt"

func main() {
	//konversi type data untuk merubah ke type data lainya
	//merubah type data int64 menjadi int8
	var age int64 = 30
	var newAge int8 = int8(age)

	fmt.Println(age)
	fmt.Println(newAge)

	//konversi type data byte menjadi string
	var name string = "Hello"
	var char = name[2]
	var charString string = string(char)

	fmt.Println(name)       // Hello
	fmt.Println(char)       // 108
	fmt.Println(charString) // l

}
