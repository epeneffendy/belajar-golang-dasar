package main

import "fmt"

func main() {
	fmt.Println("=============CLOSURE====================")
	nameClosure := "Effendy Anwar"
	printName := func() {
		//varibale nameClosure akan barubah
		// nameClosure = "Kamandaka Madava Anwar"
		fmt.Println(nameClosure)
	}

	//Closure
	//kombinasi dari function yang dibandle bersama dalam scope di sekitarnya
	//closure adalah function di dalam function
	//hati-hati dalam penggunaan closure, dapat merubah nilai di luar function
	printName()
}
