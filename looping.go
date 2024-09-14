package main

import "fmt"

func main() {

	//For Loop Statement
	count := 1

	for count <= 5 {
		fmt.Println("Perulangan ke - ", count)
		count++
	}

	//For loop dengan init statement dan post statement
	for counter := 1; counter <= 7; counter++ {
		fmt.Println("Perulangan dengan init statement ke - ", counter)
	}

	//Looping dengan range

	//Perulangan dengan array
	fruits := []string{
		"Nanas",
		"Mangga",
		"Melon",
	}

	for index, fruit := range fruits {
		fmt.Println("Index : ", index, " = ", fruit)
	}

	fmt.Println("========================Perulangan dengan map================================")
	website := map[string]string{
		"name": "Google",
		"url":  "www.google.com",
	}

	for key, value := range website {
		fmt.Println(key, " : ", value)
	}

	fmt.Println("==================BREAK & CONTINUE====================")
	//break untuk menghentikan seluruh perulangan dan perulangan di anggap selesai
	//continue untuk menghentikan perulangan saat ini dan langsung melanjutkan perulangan selanjutnya

	fmt.Println("==================BREAK====================")
	for i := 1; i <= 5; i++ {
		fmt.Println("Perulangan Break : ", i)
		if i == 3 {
			fmt.Println("Break - ", i)
			break
		}
	}

	fmt.Println("==================COUNTINUE====================")
	for x := 1; x <= 10; x++ {
		if x%2 == 0 {
			continue
		}

		fmt.Println("Index : ", x)
	}
}
