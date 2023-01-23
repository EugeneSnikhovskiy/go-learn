package main

import (
	"errors"
	"fmt"
)

func main() {
	// мы можем указать размер массива, не будет компилиться, если попытаться сослаться на неправильный индекс вне ренджа
	// также не будет компилится, если исходный массив и функция, которая его принимает аргументом имеют разный рендж
	// можно считать просто, что массивы с разной длиной - это разные типы
	messages := [3]string {"1", "2", "3"}
	fmt.Println(messages)
	fmt.Println(messages[2])
	messages[2] = "5"
	fmt.Println(messages[2])
	print(messages)
	fmt.Println(messages)

	// это уже слайс, обертка над массивом, не указываем рендж.
	// Слайс под капотом хранит в себе ссылку на массив
	// слайсы передаются по ссылке
	messages2 := []string {"9", "8", "7"}
	fmt.Println(messages2)
	print2(messages2)
	fmt.Println(messages2)

	fmt.Println("----------1------------")

	var slice []string
	// выдаст рантайм ошибку потому, что под капотом ссылка nill, решается при помощи make
	// slice[0] = "1"
	fmt.Println(slice)
	fmt.Println(len(slice))

	fmt.Println("----------2------------")

	slice2 := make([]string, 5)
	slice2[2] = "1"
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2)) // capacity равен длине массива, по дефолту, т.к. его не передали
	slice2 = append(slice2, "6") // т.к. мы не задали capacity и втулили лишний элемент, произойдет переаллокация и создастся массив с capacity в 2 раза длиннее
	fmt.Println(len(slice2))
	// начиная с определенной точки capacity увеличивается не в 2 раза (видно на длине в 10000)
	fmt.Println(cap(slice2))

	fmt.Println("----------3------------")

	slice3 := make([]string, 5, 15)
	fmt.Println(cap(slice3)) // capacity равен 15, т.е. мы можем 
}

func print(msgs [3]string) error {
	if len(msgs) == 0 {
		return errors.New("Empty")
	}
	msgs[1] = "10"
	fmt.Println(msgs) // не меняет оригинальный массив messages в main
	return nil
}


func print2(msgs []string) error {
	if len(msgs) == 0 {
		return errors.New("Empty")
	}
	msgs[1] = "10"
	fmt.Println(msgs) // меняет оригинальный слайс messages2 в main
	return nil
}