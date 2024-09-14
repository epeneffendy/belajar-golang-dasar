package main

import "fmt"

func main() {
	var x = 10
	var y = 5

	//operasi penjumlahan
	penjumlahan := x + y
	fmt.Println(penjumlahan)

	//operasi pengurangan
	pengurangan := x - y
	fmt.Println(pengurangan)

	//operasi perkalian
	perkalian := x * y
	fmt.Println(perkalian)

	//operasi pembagian
	pembagian := x / y
	fmt.Println(pembagian)

	//operasi modulus
	mod := x % y
	fmt.Println(mod)

	fmt.Println("=================Augmented Assingments==========================")
	//Augmented Assignments
	//memoersingkat operasi matematika
	//a = a + 5 => a += 5
	var a = 10
	fmt.Println(a)
	a += 5
	fmt.Println(a)

	fmt.Println("=================Unary Operator==========================")
	//mempersingkat operasi matematika dengan menaikan 1 nilai kedalam variable
	// c = c + 1 => c++
	var c = 10
	fmt.Println(c)
	c++

	fmt.Println(c)

	fmt.Println("=================Operasi Perbandingan==========================")
	name := "Effendy Anwar"
	alias := "Effendy Anwar"
	nickname := "Effendy"

	var result1 = name == alias
	fmt.Println(result1)

	var result2 = name == nickname
	fmt.Println(result2)

	var result3 = alias != nickname
	fmt.Println(result3)

	var age1 = 29
	var age2 = 30

	var result4 = age1 > age2
	fmt.Println(result4)

	var result5 = age1 < age2
	fmt.Println(result5)

	var result6 = age1 != age2
	fmt.Println(result6)
}
