package main

import "fmt"

func main() {
	printMessage("Hello world")
	printMessage(sayHello("Eugene", 34))
	// _ сбросит переменную
	var message1, _ = checkAge(18)
	var message2, _ = checkAge(46)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(checkAge(17))
}

func printMessage(message string) {
	fmt.Println(message)
}

func sayHello(name string, age int) string {
	// как конкатинировать строки и числа
	var message = fmt.Sprintf("Hello, %s! Your age is %d years", name, age)
	return message
}

func checkAge(age int) (string, bool) {
	if age >= 18 && age <= 45 {
		var message = "You are middle"
		return message, true
	} else if age > 45 {
		return "You are old", true
	}

	return "You are yong", false
}
