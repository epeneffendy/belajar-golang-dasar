package main

import "fmt"

//kumpulan data field yang dideklarasikan dengan tipe data
//struct berfungsi sebagai template untuk kumpulan beberapa data
//mirip seperti kumpulan object atau class

type User struct {
	Name, Email string
	Age         int
}

func main() {
	var user User

	//STRUCT LITERALS
	//dengan menyebutkan field
	var user1 = User{
		Name:  "Kamandaka Madava Anwar",
		Email: "kamandaka@gmail.com",
		Age:   3,
	}

	//tanpa menyebutkan field (urutan data harus sesuai dengan structnya)
	var user2 = User{
		"Indah Setyowati",
		"indahsekali13@gmail.com",
		20,
	}

	user.Name = "Effendy Anwar"
	user.Email = "epen.effendy@gmail.com"
	user.Age = 20

	fmt.Println(user)

	//untuk mengambil data berdasarkan field
	fmt.Println(user.Name)

	fmt.Println(user1)
	fmt.Println(user2)

	// karena struct user sudah ditambahkan method, sekarang kita bisa menggunakannya
	fmt.Println(user.GetName())
	fmt.Println(user.GetEmail())
}

//STRUCT METHOD
// Membuat method yang akan ditempelkan pada struct User.
// kita menuliskan struct-nya di depan nama function,
// lalu bisa kita gunakan struct tersebut dalam function yang kita buat.
func (user User) GetName() string {
	return user.Name
}

func (user User) GetEmail() string {
	return user.Email
}
