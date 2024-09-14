package main

import "fmt"

//pointer adalah tipe data yang menyimpan alamat memori dari suatu variable
//Operator pada Golang Pointer
//& (address operator) : digunakan untuk mengambil alamat memori dari suatu variabel.
//* (dereference operator) : digunakan untuk mengambil nilai yang ditunjuk oleh sebuah pointer.
//new : digunakan untuk mengalokasikan memori untuk sebuah variabel dan mengembalikan pointer ke alamat memori yang baru dialokasikan.
//make : digunakan untuk mengalokasikan memori dan menginisialisasi tipe data yang bersifat kompleks seperti slice, map, atau channel.

func main() {
	// Deklarasi variabel x sebagai integer
	x := 5

	// Deklarasi pointer p sebagai integer
	var p *int

	// Memberikan alamat dari x ke pointer p

	p = &x

	// Menampilkan nilai dari x
	fmt.Println("Nilai x :", x)

	// Menampilkan alamat memori dari x
	fmt.Println("Alamat memori x :", &x)

	// Menampilkan nilai yang ditunjuk oleh pointer p
	fmt.Println("Nilai yang ditunjuk p :", *p)

	fmt.Println("==================PAS BY VALUE====================")
	a := 5
	b := 3
	c := add(a, b)
	fmt.Println("pass by value : ", c)

	fmt.Println("==================PAS BY REFENECE====================")
	d := 5
	e := 3
	fmt.Println(d)
	sum(&d, &e)
	fmt.Println(d)

	fmt.Println("=================Pointer Pada Function Parameter==================")
	value := true
	fmt.Println(value)

	changeValueToFalse(&value)
	fmt.Println(value)
}

func add(a int, b int) int {
	return a + b
}

func sum(d *int, e *int) {
	*d = *d + *e
}

//Pointer Pada Function Parameter
// parameter pada sebuah function, secara default parameter tersebut adalah pass by value
func changeValueToFalse(value *bool) {
	*value = false
}
