package main

import "fmt"

func main() {
	//slice merupakan potongan data dari array
	//array panjangnya tidak bisa berubah slice panjangan nya bisa berubah
	//slice selalu terhubung dengan array karena slice mengakses sebagian atau seluruh array
	//slice memiliki 3 buah data yaitu pointer(sebuah petunjuk data pertama pada array) lenght(panjang data slice, tidak boleh lebih dari capacity) capacity(kapasitas dari slice)

	//MEMBUAT SLICE
	//array[low:high]	Membuat slice dimulai dari index low hingga sebelum index high
	//array[low:]	Membuat slice dimulai dari index low hingga index terakhir dari array
	//array[:high]	Membuat slice dimulai dari index 0 hingga index high
	//array[:]	Membuat slice dimulai dari index 0 hingga index terakhir dari array

	//FUNGSI SLICE
	//len(slice)	Mendapatkan panjang slice
	//cap(slice)	Mendapatkan kapasitas slice
	//append(slice, data)	Membuat slice baru dengan menambah data ke posisi terakhir slice. Jika kapasitas penuh, maka akan otomatis membuat array baru
	//make([]TypeData, length, capacity)	Membuat slice baru
	//copy(destination, source)	Menyalin slice dari source ke destination

	var fruits = [...]string{
		"Apel",
		"Jeruk",
		"Semangka",
		"Nanas",
		"Anggur",
		"Melon",
		"Pepaya",
	}

	var slice = fruits[3:5] //Semangka Nanas Anggur
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println("==================APPEND=======================")
	//append untuk menambah data ke posisi terakhir slice, jika kapasitas penuh otomatis akan membuat array baru
	chars := [...]string{"A", "B", "C", "D", "E", "F"}
	charSlice1 := chars[4:]
	fmt.Println(chars)      //A B C D E F
	fmt.Println(charSlice1) //E F

	charSlice2 := append(charSlice1, "G")
	fmt.Println(charSlice2) //E F G
	charSlice2[0] = "E Lagi"
	fmt.Println(charSlice2) //E Lagi F G
	fmt.Println(chars)

	fmt.Println("==================MAKE=======================")
	//make digunakan untuk membuat slice dari awal tanpa menggunakan array
	days := make([]string, 2, 6)
	days[0] = "Senin"
	days[1] = "Selasa"

	fmt.Println(days)

	fmt.Println("==================COPY SLICE=======================")
	copyDays := make([]string, len(days), cap(days))
	copy(copyDays, days)
	fmt.Println(copyDays)

	fmt.Println("==================MAP=======================")
	//map adalah type data key value-pairs, digunakan untuk mengambil nilai berdasarkan key
	//len(map) :	Menghitung jumlah data yang ada di dalam map
	//map[key] :	Mengambil data berdasarkan kata kunci key
	//map[key] = value :	Menambah atau merubah data di map
	//make(map[TypeKey]TypeValue) :	Membuat map baru
	//delete(map, key) :	Menghapus data di map berdasarkan kata kunci key

	website := map[string]string{
		"name":   "Search Engine",
		"domain": "google.com",
	}

	fmt.Println(website)

	//mengakses data pada map
	fmt.Println(website["name"])
	fmt.Println(website["domain"])

	//menambah data pada map
	website["tagline"] = "Search Engine Populer"

	//merubah data pada map
	website["domain"] = "www.google.com"

	fmt.Println(website["tagline"])
	fmt.Println(website["domain"])

	//menghapus data pada map
	delete(website, "tagline")
	fmt.Println(website)

	fmt.Println("==================MAKE MAP=======================")
	month := make(map[string]string)

	month["indonesia"] = "Januari"
	month["english"] = "January"

	fmt.Println(month)

}
