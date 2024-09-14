package main

import "fmt"

func main() {
	age := 3
	//If Expression
	if age == 22 {
		fmt.Println(age)
	}

	//else expression
	if age > 17 {
		fmt.Println("Sudah cukup umur")
	} else {
		fmt.Println("Belum cukup umur")
	}

	//else if expression
	if age > 17 {
		fmt.Println("Sudah cukup umur")
	} else if age > 5 && age < 17 {
		fmt.Println("Belum cukup umur")
	} else {
		fmt.Println("Masih balita")
	}

	//Short statement pada if expression
	if count := 100; count > 50 {
		fmt.Println("Count lebih dari 50")
	}

	//Switch Expression
	var category = "CAT1"

	switch category {
	case "CAT1":
		fmt.Println("Kategori CAT1")
	case "CAT2":
		fmt.Println("Kategori CAT2")
	default:
		fmt.Println("Katagori tidak dikenali")
	}

	//Short Statement Switch Expression
	switch married := true; married {
	case true:
		fmt.Println("Sudah Menikah")
	case false:
		fmt.Println("Belum Menikah")
	}

	//Switch tanpa kondisi
	id := "A1"
	switch {
	case id == "A1":
		fmt.Println("ID : A1")
	case id == "A2":
		fmt.Println("ID : A2")
	default:
		fmt.Println("ID Invalid")
	}

}
