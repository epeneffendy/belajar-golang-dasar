package main

import "fmt"

// defer
// statement yang menunda pengeksekusian sebuah function
//defer statement akan selalu di jalankan meskipun terjadi error

//panic
//function untuk menghentikan program
//biasanya di panggil saat terjadi error

//recover
//function untuk menangkap data dari panic function
//kegunaan seperti catch

func main() {
	//kode ini akan tereksekusi mestkipun memanggi panic function
	// defer hello()
	// fmt.Println("Show Me!")

	// panic("Panic dipanggil")

	// sedangkan kode ini tidak akan dieksekusi, karena program berhenti
	// setelah panic function dipanggil
	// fmt.Println("tidak akan dieksekusi")

	//contoh meskipun terjadi error
	// div := 0
	// result := 3 / div
	// fmt.Println(result)

	defer selesai()
	checkSudahMakan(false)
}

func hello() {
	fmt.Println("Hello World")
}

func checkSudahMakan(sudahMakan bool) {
	if sudahMakan == false {
		panic("Ups!, kamu belum makang!")
	}

	fmt.Println("Mantab, kamu sudah makan!")
}

func selesai() {
	er := recover()
	if er != nil {
		fmt.Println("Ada error", er)
	}

	fmt.Println("Selesai")
}
