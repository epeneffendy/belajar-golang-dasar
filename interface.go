package main

import "fmt"

//interface sebuah tipe data yang tidak memiliki implementasi secara langsung (abstract)
//terdapat definis function/method

//membuat interfaec
type Engine interface {
	ShowName() string
	Start() string
	Stop() string
}

func main() {
	// Mendeklarasikan variabel engine1 bertipe struct MyEngine
	var engine1 MyEngine
	engine1.Name = "V8"

	// karena MyEngine sudah memiliki semua method/function yang sama dengan
	// interface Engine, maka secara otomatis telah mengimplementasikan interface tersebut dan
	// bisa kita gunakan untuk mengisi parameter pada function CheckEngine
	CheckEngine(engine1)

	//interface kosong
	fmt.Println("=====================INTERFACE KOSONG==================")
	fmt.Println(ReturnAny(1))
	fmt.Println(ReturnAny(2))
	fmt.Println(ReturnAny(3))
}

//membuat function untuk menerima parameter engine
func CheckEngine(engine Engine) {
	fmt.Println("Engine Name : ", engine.ShowName())
	fmt.Println("Check Strat Engine : ", engine.Start())
	fmt.Println("Check Stop Engine : ", engine.Stop())
}

//Membuat struct
type MyEngine struct {
	Name string
}

// Membuat struct method untuk ShowName
func (myEngine MyEngine) ShowName() string {
	return myEngine.Name
}

func (MyEngine MyEngine) Start() string {
	return "Engine Start......"
}

func (myEngine MyEngine) Stop() string {
	return "Engine Stop......."
}

//interface kosong
func ReturnAny(i int) interface{} {
	if i == 1 {
		return "string"
	} else if i == 2 {
		return 0
	} else {
		return true
	}
}
