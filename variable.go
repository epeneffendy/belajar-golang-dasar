package main

import "fmt"

func main() {
	// variable di deklarasikan dengan kata kunci var dan diikuti type data nya

	var name string //contoh deklarasi variable tyep data string
	//variable name diisi nilai dengan tanda petik untuk string
	name = "Ruang Developer"

	fmt.Println(name)

	//mengubah nilai data dalam variable dengan syarat harus memiliki type data yang sama
	// nilai name yang awalnya bernilai "Ruang Developer" berubah menjadi "Belajar Golang"
	name = "Belajar Golang"
	fmt.Println(name)

	//mendeklarasikan variable tanpa kata kunci va
	firstName := "Golang Keren"
	fmt.Println(firstName)

	//mendeklarasikan banyak variable secara bersamaan
	var (
		first_name = "Effendy"
		last_name  = "Anwar"
	)

	fmt.Println(first_name)
	fmt.Println(last_name)

	//Constant merupakan varibale yang nilainya tidak dapat dirubah
	//jika dipaksa untuk di rubah akan terjadi error
	const web = "www.google.com"
	fmt.Println(web)

}
