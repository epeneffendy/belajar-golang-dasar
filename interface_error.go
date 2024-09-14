package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	result, err := Uppercase("")

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err.Error())
	}
}

func Uppercase(text string) (string, error) {
	if text == "" {
		return "", errors.New("Text tidak boleh kosong!")
	} else {
		return strings.ToUpper(text), nil
	}
}
