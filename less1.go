package main

import (
	"fmt"
	"reflect" // либа аналог typeof в javascript
)

var a, b, c int;

func main() {
    a, b, c := 1, 2, 3;
    a, b = b, a; // поменять значение переменных местами
    a, _, c = 5, 6, 7; // _ заигнорим множественное присвоение

    var intVar int = 2;
    var float1Var float32 = 22.55;
	var a1 int = 2;
	const msg string = "Hello";
	var msg2 string = "HI";
	message := "Hello world";
    fmt.Println(float1Var, intVar, msg, msg2, message, a1);
    fmt.Println(reflect.TypeOf(message));
    fmt.Println(a, b, c);
    print();
}

func print() {
    // сошлется на глобальные переменные и выведет начальные значения (нули)
    fmt.Println(a, b, c);
}