package main

import "fmt"

func main() {
	//cara membuat array dengan menambahkan jumlah data sebelum type data
	var fruits [4]string

	//isi array dengan value
	fruits[0] = "Jambu"
	fruits[1] = "Apel"
	fruits[2] = "Jeruk"
	fruits[3] = "Semangka"

	//mengambil data tertentu pada array
	var apel = fruits[1]
	fmt.Println(apel)

	//deklarasi array secara langsung
	//array[index] : mengambil nilai array berdasarkan index
	//array[index] = value : mengisi data array berdasarkan posisi index
	//len(array) : menghitaung panjang array
	var numbers = [4]int{1, 2, 3, 4}
	fmt.Println(numbers)
	fmt.Println(numbers[2])
	fmt.Println(len(numbers))
}
