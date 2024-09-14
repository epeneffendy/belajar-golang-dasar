package main

import "fmt"

func main() {
	//type data Integer
	//Signed Integer
	// => type data yang dapat diisi bilangan negatif
	//contoh => int8, int16, int32, int64

	//Unsigned Integer
	// => type data yang tidak dapat digunakan bilangan negatif
	//contoh => uint8, uint16, uint32, unint64

	//Type data Floating Point
	// => type data yang bisa diisi bilangan desimal
	// contoh => float32, float64, complax64, complax128
	//untuk type data complax64 dan complax128 untuk perhitungan matematika yang presisi

	fmt.Println("====================INTEGER DAN FLOATING========================")
	fmt.Println(21)   //integer
	fmt.Println(21.5) //floating

	fmt.Println("====================BOOLEAN========================")
	//Type data Boolean
	// type data yang bernilai true dan false (benar atau salah)
	fmt.Println(true)
	fmt.Println(false)

	fmt.Println("====================STRING========================")
	//Type data String
	fmt.Println("Saya sedang belajar Golang!")

	fmt.Println("====================MENGHITUNG PANJANG KARAKTER========================")
	// fungsi len untuk mengitunng jumlah karakter termasuk spasi
	fmt.Println(len("Ruang Developer"))
	// mengampil karakter di posisi tertentu
	// data kembalian yang di berikan merupakan data byte. value 108 adalah reperesentasi dari byte huruf l
	fmt.Println("Hello"[2]) // 108

}
