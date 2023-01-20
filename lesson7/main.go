package main

import (
	"fmt"
)

var message string
// стандартная функция, срабатывает при инициализации пакетов
// всегда срабатывает ПЕРЕД вызовом main
func init() {
	message = "hello"
}

func main() {
	fmt.Println(message)
	print(message)
	fmt.Println(message) // будет старое значение, не измененное в print
	print2(&message) // referencing
	fmt.Println(message)

	var num = 5
	var p *int // инициализируем указатель
	fmt.Println(p) // nil т.к. пустой
	p = &num
	fmt.Println(p) // ссылка на память, т.к. выше указали, куда ссылаться
	fmt.Println(num) // 5
	*p = 10
	fmt.Println(num) // 10 т.к. изменили значение переменной по ссылке
}

func print(message string) { // тут под message выделяется своя новая переменная
	message += " world"
	fmt.Println(message)
}

func print2(message *string) { // тут message передается по ссылке и будет менять оригинальное значение
	*message += " ref world" // dereferencing
	fmt.Println(*message)
}