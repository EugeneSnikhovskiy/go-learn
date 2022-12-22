package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	var message, error = checkAge(17)
	if error != nil {
		log.Fatal(error)
	} else {
		fmt.Println(message)
	}
}

func checkAge(age int) (string, error) {
	if age >= 18 {
		return "Ok", nil
	}
	// best practice текст ошибки с маленькой буквы
	// best practice ошибка всегда возвращается последним аргументом
	return "Boy", errors.New("error text you are too young")
}
