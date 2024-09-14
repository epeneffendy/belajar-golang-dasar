package main

import "fmt"

//tipe data kosong (tidak bernilai)
//nil hanya bisa digunakan bebrapa tipe data, interface, map, slice, pointer, function, channel
func main() {
	user := NewMap("")

	if user == nil {
		fmt.Println("Kosong")
	} else {
		fmt.Println(user)
	}
}

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{"Name ": name}
}
