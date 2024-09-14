package main

import (
	"fmt"
	"strings"
)

// Type Declaration function as parameter
type Formatter func(string) string
type Authorize func(string) bool

func main() {
	fmt.Println("=============FUNCTION====================")
	helloWorld()

	fmt.Println("=============FUNCTION PARAMETER====================")
	sayHello("Effendy", "Anwar")

	fmt.Println("=============FUNCTION RETURN VALUE====================")
	result := sum(10, 5)
	fmt.Println("Hasil Penjumlahan : ", result)

	isMarried := isMarried(true)
	fmt.Println(isMarried)

	fmt.Println("=============FUNCTION RETURN MULTIPLE VALUE====================")
	area, perimeter := rectangle(5)
	fmt.Println("Luas Persegi :", area)
	fmt.Println("Keliling Persegi :", perimeter)

	fmt.Println("=============VARIADIC FUNCTION====================")
	sumResult := sumNumbers(5, 5, 5, 5)
	fmt.Println(sumResult)

	fmt.Println("=============VARIADIC FUNCTION SLICE PARAMETER====================")
	numberSlice := []int{10, 10, 10}
	sliceSumResult := sliceSumNumber(numberSlice...)
	fmt.Println(sliceSumResult)

	fmt.Println("=============FUNCTION AS VALUE====================")
	hi := sayHiTo
	resultHiTo := hi("Effendy Anwar")

	fmt.Println(resultHiTo)

	fmt.Println("=============FUNCTION AS PARAMETER====================")
	formatAndPrint("hello world", upper)
	formatAndPrint("HELLO WORLD", lower)

	formatAndPrinted("effendy anwar", upper)
	formatAndPrinted("EFFENDY ANWAR", lower)

	fmt.Println("=============ANONYMOUS FUNCTION====================")
	//contoh 1 - dengan variable
	onlyAdmin := func(username string) bool {
		return username == "admin"
	}

	authorized("admin", onlyAdmin)

	//contoh 2 - langsung di parameter sebagai argument
	authorized("admin", func(username string) bool {
		return username == "member"
	})

	fmt.Println("=============RECURSIVE FUNCTION====================")
	resultRecursive := doFactorial(5)
	fmt.Println(resultRecursive)

	resultNonRecursive := doFactorialWithForLoop(5)
	fmt.Println(resultNonRecursive)

}

func helloWorld() {
	fmt.Println("Hello World from function")
}

// Function Parameter
func sayHello(firstName string, lastName string) {
	fmt.Println("Hello ", firstName, lastName)
}

// Function Return Value
func sum(value1 float32, value2 float32) float32 {
	return value1 + value2
}

func isMarried(married bool) string {
	if married {
		return "Sudah Menikah"
	} else {
		return "Belum Menikah"
	}
}

// Function Multiple return
func rectangle(sideLength float32) (float32, float32) {
	area := sideLength * sideLength
	perimeter := sideLength * 4

	return area, perimeter
}

// Vaiadic Function
// function yang dapat menerima sejumlah variable sebagai argumen
// argument disebut sebagai Varargs
// Varagrs berbeda dengan tipe data array
func sumNumbers(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

//SALAH
//func sum(numbs ...int, name string)
//BENAR
//func sum(name string, numbs ...int)

// Slice sebagai parameter
func sliceSumNumber(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}

	return result
}

// function as value
// function juga merupakan tipe data dan bisa dianggap nilai
// menyimpan function kedalam variable
func sayHiTo(name string) string {
	return "Hi, " + name
}

// function as parameter
func formatAndPrint(text string, formatter func(string) string) {
	formatted := formatter(text)
	fmt.Println(formatted)
}

func formatAndPrinted(text string, formatter Formatter) {
	formatted := formatter(text)
	fmt.Println(formatted)
}

func upper(text string) string {
	return strings.ToUpper(text)
}

func lower(text string) string {
	return strings.ToLower(text)
}

// Anonymous Function
// funtion yang tidak memiliki nama
func authorized(username string, authorize Authorize) {
	if authorize(username) {
		fmt.Println("Authorize")
	} else {
		fmt.Println("Unauthorize")
	}
}

// function recursive
// function yang memanggil/mengeksekusi dirinya sendiri
// function yang bisa digunakan untuk perulangan
// kelebihan : kode mudah ditulis, kurangi pemanggilan yang tidak perlu
// kekurangan : lebih lambat, terkadang memerlukan memori yang lebih
func doFactorial(value int) int {
	if value == 1 {
		return value
	}

	return value * doFactorial(value-1)
}

func doFactorialWithForLoop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}
