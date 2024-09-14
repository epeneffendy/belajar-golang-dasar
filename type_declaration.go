package main

import "fmt"

func main() {
	//type declaration untuk membuat ulang type data baru dari type data yang sudah ada
	//menggunakan kata kunci type lalu diikuti dengan type data baru kemudian type data aslinya

	type PhoneNumber string
	type Address string
	type Age int

	var phoneNumber PhoneNumber = "089900000000"
	var address Address = "Malang"
	var age Age = 30

	fmt.Println(phoneNumber)
	fmt.Println(address)
	fmt.Println(age)

}
